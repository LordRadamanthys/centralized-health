package output

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type TrainingPort interface {
	GetTrainingByUser(string) (*domain.TrainingDomain, *rest_errors.RestErr)
	CreateTraining(*domain.TrainingDomain) *rest_errors.RestErr
	UpdateTraining(*domain.TrainingDomain) *rest_errors.RestErr
}
