package output

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type VaccinesPort interface {
	GetVaccineByUser(string) (*domain.VaccinesDomain, *rest_errors.RestErr)
	CreateVaccine(*domain.VaccinesDomain) (*domain.VaccinesDomain, *rest_errors.RestErr)
}
