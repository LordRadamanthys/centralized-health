package service

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/output"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type userService struct {
	userRepository output.UserPort
}

func NewUserService(repository output.UserPort) *userService {
	return &userService{
		userRepository: repository,
	}
}

func GetUserByID(string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func GetUserByEmail(string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func UpdateUserByID(*domain.UserDomain) *rest_errors.RestErr {
	return nil
}

func UpdateUserByEmail(*domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func CreateUser(*domain.UserDomain) *rest_errors.RestErr {
	return nil
}

func LoginService(email string, password string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}
