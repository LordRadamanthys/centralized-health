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

func (us *userService) GetUserByID(id string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (us *userService) GetUserByEmail(email string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (us *userService) UpdateUserByID(user *domain.UserDomain) *rest_errors.RestErr {
	return nil
}

func (us *userService) UpdateUserByEmail(user *domain.UserDomain) *rest_errors.RestErr {
	return nil
}

func (us *userService) CreateUser(user *domain.UserDomain) *rest_errors.RestErr {
	return nil
}

func (us *userService) LoginService(email string, password string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}
