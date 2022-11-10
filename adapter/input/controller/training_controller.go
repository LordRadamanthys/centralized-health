package controller

import (
	"net/http"

	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/port/input"
	jwtconfig "github.com/LordRadamanthys/centralized-health/configuration/jwt_config"
	"github.com/gin-gonic/gin"
)

type trainingController struct {
	trainingService input.TrainingUseCase
}

func NewTrainingService(service input.TrainingUseCase) *trainingController {
	return &trainingController{
		trainingService: service,
	}
}

func (tc *trainingController) CreateTraining(ctx *gin.Context) {
	var trainingRequest *requests.TrainingRequest

	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	if err := ctx.ShouldBindJSON(&trainingRequest); err != nil {
		ctx.JSON(400, "invalid object")
		return
	}

	if err := tc.trainingService.CreateTraining(id, trainingRequest); err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (tc *trainingController) GetTraining(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	training, err := tc.trainingService.GetTrainingByUser(id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	ctx.JSON(http.StatusOK, training)
}

func (tc *trainingController) UpdateTraining(ctx *gin.Context) {
	var trainingRequest *requests.TrainingRequest

	day_week := ctx.Param("day")
	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	if err := ctx.ShouldBindJSON(&trainingRequest); err != nil {
		ctx.JSON(400, "invalid object")
		return
	}

	if err := tc.trainingService.UpdateTraining(id, day_week, trainingRequest); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
