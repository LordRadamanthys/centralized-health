package input

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type TrainingUseCase interface {
	GetTrainingByUser(string) (*domain.TrainingDomain, *rest_errors.RestErr)
	CreateTraining(*domain.TrainingDomain) *rest_errors.RestErr
}
