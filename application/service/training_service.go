package service

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/converter"
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
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

func (ts *trainingService) UpdateTraining(userID string, day_week string, request *requests.TrainingDetails) *rest_errors.RestErr {

	trainingDomain, err := ts.GetTrainingByUser(userID)
	if err != nil {
		return err
	}

	details := &domain.TrainingDetails{
		Muscle:      request.Muscle,
		Activity:    request.Activity,
		Series:      request.Series,
		Repetitions: request.Repetitions,
	}

	switch day_week {
	case "seg":
		trainingDomain.Seg = []domain.TrainingDetails{*details}
	case "ter":
		trainingDomain.Ter = []domain.TrainingDetails{*details}
	case "qua":
		trainingDomain.Qua = []domain.TrainingDetails{*details}
	case "qui":
		trainingDomain.Qui = []domain.TrainingDetails{*details}
	case "sex":
		trainingDomain.Sex = []domain.TrainingDetails{*details}
	case "sab":
		trainingDomain.Sab = []domain.TrainingDetails{*details}
	case "dom":
		trainingDomain.Dom = []domain.TrainingDetails{*details}
	}

	return ts.trainingRepository.UpdateTraining(trainingDomain)
}
