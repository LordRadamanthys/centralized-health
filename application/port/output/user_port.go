package output

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPort interface {
	GetUserByID(string) (*domain.UserDomain, *rest_errors.RestErr)
	GetUserByEmail(string) (*domain.UserDomain, *rest_errors.RestErr)
	UpdateUserByID(primitive.ObjectID, *requests.UserRequest) *rest_errors.RestErr
	UpdateUserByEmail(*requests.UserRequest) (*domain.UserDomain, *rest_errors.RestErr)
	CreateUser(*requests.UserRequest) *rest_errors.RestErr
}
