package input

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type TrainingUseCase interface {
	GetTrainingByUser(string) (*domain.TrainingDomain, *rest_errors.RestErr)
	CreateTraining(string, *requests.TrainingRequest) *rest_errors.RestErr
	UpdateTraining(string, string, *requests.TrainingRequest) *rest_errors.RestErr
}
