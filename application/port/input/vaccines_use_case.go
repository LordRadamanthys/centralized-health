package input

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type VaccinesUseCase interface {
	GetVaccineByUser(string) ([]*domain.VaccinesDomain, *rest_errors.RestErr)
	GetVaccineByID(vaccineID string) (*domain.VaccinesDomain, *rest_errors.RestErr)
	CreateVaccine(string, *domain.VaccinesDomain) *rest_errors.RestErr
	InsertVaccineDocument(userID string, idVaccine string, fileName string, bufferFile []byte) *rest_errors.RestErr
}
