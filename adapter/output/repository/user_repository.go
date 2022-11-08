package repository

import (
	"context"

	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/database/mongodb"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (*userRepository) GetUserByID(string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByEmail(email string) (*domain.UserDomain, *rest_errors.RestErr) {
	var user domain.UserDomain
	result := mongodb.MongoConnection.Collection("users").FindOne(context.TODO(), bson.M{"email": email})

	if err := result.Decode(&user); err != nil {
		return nil, rest_errors.NewBadRequestError("invalid user")
	}

	return &user, nil
}

func (*userRepository) UpdateUserByID(id primitive.ObjectID, obj *requests.UserRequest) *rest_errors.RestErr {
	return nil
}

func (*userRepository) UpdateUserByEmail(*requests.UserRequest) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) CreateUser(user *requests.UserRequest) *rest_errors.RestErr {
	_, err := mongodb.MongoConnection.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}
