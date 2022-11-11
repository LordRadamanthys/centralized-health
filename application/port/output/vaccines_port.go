package output

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VaccinesPort interface {
	GetVaccineByUser(primitive.ObjectID) ([]*domain.VaccinesDomain, *rest_errors.RestErr)
	CreateVaccine(*domain.VaccinesDomain) *rest_errors.RestErr
	GetVaccineById(idVaccine primitive.ObjectID) (*domain.VaccinesDomain, *rest_errors.RestErr)
	UpdateVaccine(idVaccine primitive.ObjectID, obj *domain.VaccinesDomain) *rest_errors.RestErr
}
