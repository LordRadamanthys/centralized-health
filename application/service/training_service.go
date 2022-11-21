package service

import (
	"strings"

	"github.com/LordRadamanthys/centralized-health/adapter/input/converter"
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/common"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/output"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type trainingService struct {
	trainingRepository output.TrainingPort
}

func NewTrainingService(repository output.TrainingPort) *trainingService {
	return &trainingService{
		trainingRepository: repository,
	}
}

func (ts *trainingService) GetTrainingByUser(userID string) (*domain.TrainingDomain, *rest_errors.RestErr) {
	return ts.trainingRepository.GetTrainingByUser(userID)
}

func (ts *trainingService) CreateTraining(userID string, obj *requests.TrainingRequest) *rest_errors.RestErr {
	idConv, errConv := primitive.ObjectIDFromHex(userID)
	if errConv != nil {
		return rest_errors.NewBadRequestError("invalid id")
	}

	domain := converter.ConverterRequestIntoDomainTraining(obj)
	domain.IdUser = idConv
	return ts.trainingRepository.CreateTraining(domain)
}

func (ts *trainingService) UpdateTraining(userID string, day_week string, request *requests.TrainingRequest) *rest_errors.RestErr {

	if day_week == "" {
		return rest_errors.NewBadRequestError("day week cannot be empty")
	}

	trainingDomain, err := ts.GetTrainingByUser(userID)
	if err != nil {
		return err
	}

	switch strings.ToLower(day_week) {
	case common.Segunda:
		trainingDomain.Seg = request.Training
	case common.Terca:
		trainingDomain.Ter = request.Training
	case common.Quarta:
		trainingDomain.Qua = request.Training
	case common.Quinta:
		trainingDomain.Qui = request.Training
	case common.Sexta:
		trainingDomain.Sex = request.Training
	case common.Sabado:
		trainingDomain.Sab = request.Training
	case common.Domingo:
		trainingDomain.Dom = request.Training
	default:
		return rest_errors.NewBadRequestError("invalid day week")
	}

	return ts.trainingRepository.UpdateTraining(trainingDomain)
}
