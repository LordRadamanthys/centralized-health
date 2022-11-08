package output

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type UserPort interface {
	GetUserByID(string) (*domain.UserDomain, *rest_errors.RestErr)
	GetUserByEmail(string) (*domain.UserDomain, *rest_errors.RestErr)
	UpdateUserByID(*domain.UserDomain) *rest_errors.RestErr
	UpdateUserByEmail(*domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
	CreateUser(*requests.UserRequest) *rest_errors.RestErr
}
