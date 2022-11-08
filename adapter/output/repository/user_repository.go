package repository

import (
	"context"

	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/database/mongodb"
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

func (*userRepository) CreateUser(user *requests.UserRequest) *rest_errors.RestErr {
	_, err := mongodb.MongoConnection.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}
