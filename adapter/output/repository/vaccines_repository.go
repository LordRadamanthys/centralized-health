package repository

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/database/mongodb"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type vaccinesRepository struct{}

const collection_vaccines_name = "vaccines"

func NewVaccinesRepository() *vaccinesRepository {
	return &vaccinesRepository{}
}

func (*vaccinesRepository) GetVaccineByUser(id primitive.ObjectID) ([]*domain.VaccinesDomain, *rest_errors.RestErr) {

	result, err := mongodb.MongoConnection.Collection(collection_vaccines_name).Find(context.TODO(), bson.M{"iduser": id})

	if err != nil {
		return nil, rest_errors.NewBadRequestError(err.Error())
	}

	var listVaccines []*domain.VaccinesDomain

	if err := result.All(context.TODO(), &listVaccines); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}

	fmt.Println(listVaccines)
	return listVaccines, nil
}

func (*vaccinesRepository) GetVaccineById(idVaccine primitive.ObjectID) (*domain.VaccinesDomain, *rest_errors.RestErr) {
	var vaccineDomain *domain.VaccinesDomain
	result := mongodb.MongoConnection.Collection(collection_vaccines_name).FindOne(context.TODO(), bson.M{"_id": idVaccine})

	if err := result.Decode(&vaccineDomain); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}

	return vaccineDomain, nil
}

func (*vaccinesRepository) CreateVaccine(obj *domain.VaccinesDomain) *rest_errors.RestErr {
	_, err := mongodb.MongoConnection.Collection(collection_vaccines_name).InsertOne(context.TODO(), obj)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}

func (*vaccinesRepository) UpdateVaccine(idVaccine primitive.ObjectID, obj *domain.VaccinesDomain) *rest_errors.RestErr {
	filter := bson.M{"_id": idVaccine}
	update := bson.M{"$set": obj}

	_, err := mongodb.MongoConnection.Collection(collection_vaccines_name).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}
