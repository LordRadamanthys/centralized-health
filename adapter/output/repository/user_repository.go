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

const collection_users_name = "users"

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (*userRepository) GetUserByID(id string) (*domain.UserDomain, *rest_errors.RestErr) {
	var user domain.UserDomain
	idConv, errConv := primitive.ObjectIDFromHex(id)
	if errConv != nil {
		return nil, rest_errors.NewInternalServerError("error to compare id", errConv)
	}
	result := mongodb.MongoConnection.Collection(collection_users_name).FindOne(context.TODO(), bson.M{"_id": idConv})

	if err := result.Decode(&user); err != nil {
		return nil, rest_errors.NewBadRequestError("invalid user")
	}

	return &user, nil
}

func (*userRepository) GetUserByEmail(email string) (*domain.UserDomain, *rest_errors.RestErr) {
	var user domain.UserDomain
	result := mongodb.MongoConnection.Collection(collection_users_name).FindOne(context.TODO(), bson.M{"email": email})

	if err := result.Decode(&user); err != nil {
		return nil, rest_errors.NewBadRequestError("invalid user")
	}

	return &user, nil
}

func (*userRepository) UpdateUserByID(id primitive.ObjectID, obj *domain.UserDomain) *rest_errors.RestErr {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": obj}

	_, err := mongodb.MongoConnection.Collection(collection_users_name).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (*userRepository) UpdateUserByEmail(obj *domain.UserDomain) *rest_errors.RestErr {
	filter := bson.M{"email": obj.Email}
	update := bson.M{"$set": obj}

	_, err := mongodb.MongoConnection.Collection(collection_users_name).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (*userRepository) CreateUser(user *requests.UserRequest) *rest_errors.RestErr {
	_, err := mongodb.MongoConnection.Collection(collection_users_name).InsertOne(context.TODO(), user)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}
