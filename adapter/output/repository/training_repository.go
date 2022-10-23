package repository

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type trainingRepository struct{}

func NewTrainingRepository() *trainingRepository {
	return &trainingRepository{}
}

func (*trainingRepository) GetTrainingByUser(string) (*domain.TrainingDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*trainingRepository) CreateTraining(*domain.TrainingDomain) (*domain.TrainingDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*trainingRepository) UpdateTraining(*domain.TrainingDomain) (*domain.TrainingDomain, *rest_errors.RestErr) {
	return nil, nil
}
