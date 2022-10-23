package repository

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type vaccinesRepository struct{}

func NewVaccinesRepository() *vaccinesRepository {
	return &vaccinesRepository{}
}

func (*vaccinesRepository) GetVaccineByUser(string) (*domain.VaccinesDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*vaccinesRepository) CreateVaccine(*domain.VaccinesDomain) (*domain.VaccinesDomain, *rest_errors.RestErr) {
	return nil, nil
}
