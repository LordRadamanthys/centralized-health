package controller

import (
	"net/http"

	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/input"
	jwtconfig "github.com/LordRadamanthys/centralized-health/configuration/jwt_config"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"github.com/gin-gonic/gin"
)

type ExamsController struct {
	examsController input.ExamsUseCase
}

func NewExamsService(service input.ExamsUseCase) *ExamsController {
	return &ExamsController{
		examsController: service,
	}
}

func (ec *ExamsController) GetExamsByUser(ctx *gin.Context) {

	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	response, errResponse := ec.examsController.GetExamsByUserID(id)
	if errResponse != nil {
		ctx.JSON(errResponse.Code, errResponse)
		return
	}
	ctx.JSON(http.StatusOK, response)

}

func (ec *ExamsController) CreateExam(ctx *gin.Context) {
	var examDomain *domain.ExamsDomain
	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	if err := ctx.ShouldBindJSON(&examDomain); err != nil {
		restErr := rest_errors.Cause{
			Field:   "object request",
			Message: "invalid object",
		}
		ctx.JSON(http.StatusBadRequest, restErr)
		return
	}

	if err := ec.examsController.CreateExam(id, examDomain); err != nil {
		ctx.JSON(err.Code, err)
		return

	}

	ctx.JSON(http.StatusCreated, nil)

}

func (ec *ExamsController) InsertDocument(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	idExam := ctx.Param("idExam")
	id, err := jwtconfig.NewJWTUtils().GetId(header)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	fileHeader, _ := ctx.FormFile("file")
	f, _ := fileHeader.Open()
	var size int64 = fileHeader.Size

	buffer := make([]byte, size)
	f.Read(buffer)

	if err := ec.examsController.InsertExamDocument(id, idExam, fileHeader.Filename, buffer); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
