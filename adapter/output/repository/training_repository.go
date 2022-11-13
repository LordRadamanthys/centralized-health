package repository

import (
	"context"

	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/database/mongodb"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collection_training_name = "training"

type trainingRepository struct{}

func NewTrainingRepository() *trainingRepository {
	return &trainingRepository{}
}

func (*trainingRepository) GetTrainingByUser(id string) (*domain.TrainingDomain, *rest_errors.RestErr) {
	var training domain.TrainingDomain
	idConv, errConv := primitive.ObjectIDFromHex(id)
	if errConv != nil {
		return nil, rest_errors.NewInternalServerError("error to compare id", errConv)
	}

	result := mongodb.MongoConnection.Collection(collection_training_name).FindOne(context.TODO(), bson.M{"id_user": idConv})
	if err := result.Decode(&training); err != nil {
		return nil, rest_errors.NewBadRequestError("error to find training")
	}

	return &training, nil
}

func (*trainingRepository) CreateTraining(obj *domain.TrainingDomain) *rest_errors.RestErr {

	_, err := mongodb.MongoConnection.Collection(collection_training_name).InsertOne(context.TODO(), obj)
	if err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (*trainingRepository) UpdateTraining(obj *domain.TrainingDomain) *rest_errors.RestErr {
	filter := bson.M{"id_user": obj.IdUser}
	update := bson.M{"$set": obj}

	_, err := mongodb.MongoConnection.Collection(collection_training_name).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return rest_errors.NewBadRequestError("erro to update training")
	}
	return nil
}
