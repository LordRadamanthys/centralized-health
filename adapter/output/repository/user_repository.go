package repository

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (*userRepository) GetUserByID(string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByEmail(string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) UpdateUserByID(*domain.UserDomain) *rest_errors.RestErr {
	return nil
}

func (*userRepository) UpdateUserByEmail(*domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) CreateUser(*domain.UserDomain) *rest_errors.RestErr {
	return nil
}
