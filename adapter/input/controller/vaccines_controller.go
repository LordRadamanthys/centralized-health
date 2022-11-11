package controller

import (
	"net/http"

	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/input"
	jwtconfig "github.com/LordRadamanthys/centralized-health/configuration/jwt_config"
	"github.com/gin-gonic/gin"
)

type VaccinesController struct {
	service input.VaccinesUseCase
}

func NewVaccinesController(service input.VaccinesUseCase) *VaccinesController {
	return &VaccinesController{
		service: service,
	}
}

func (vc *VaccinesController) GetVaccineByUser(ctx *gin.Context) {

	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	response, err := vc.service.GetVaccineByUser(id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (vc *VaccinesController) CreateVaccine(ctx *gin.Context) {
	var vaccine *domain.VaccinesDomain
	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	if err := ctx.ShouldBindJSON(&vaccine); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	if err := vc.service.CreateVaccine(id, vaccine); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (vc *VaccinesController) InsertDocument(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	idVaccine := ctx.Param("idVaccine")
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

	if err := vc.service.InsertVaccineDocument(id, idVaccine, fileHeader.Filename, buffer); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
