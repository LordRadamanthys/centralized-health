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

type examsRepository struct{}

const collection_exams_name = "exams"

func NewExamsRepository() *examsRepository {
	return &examsRepository{}
}

func (*examsRepository) GetExamsByUserID(id primitive.ObjectID) ([]*domain.ExamsDomain, *rest_errors.RestErr) {

	result, err := mongodb.MongoConnection.Collection(collection_exams_name).Find(context.TODO(), bson.M{"iduser": id})

	if err != nil {
		return nil, rest_errors.NewBadRequestError(err.Error())
	}

	var listExams []*domain.ExamsDomain

	if err := result.All(context.TODO(), &listExams); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}

	fmt.Println(listExams)
	return listExams, nil

}

func (*examsRepository) GetExamByID(id primitive.ObjectID) (*domain.ExamsDomain, *rest_errors.RestErr) {
	var examsDomain *domain.ExamsDomain
	result := mongodb.MongoConnection.Collection(collection_exams_name).FindOne(context.TODO(), bson.M{"_id": id})

	if err := result.Decode(&examsDomain); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}

	return examsDomain, nil
}

func (*examsRepository) CreateExam(obj *domain.ExamsDomain) *rest_errors.RestErr {
	_, err := mongodb.MongoConnection.Collection(collection_exams_name).InsertOne(context.TODO(), obj)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}

func (*examsRepository) UpdateExam(idExam primitive.ObjectID, obj *domain.ExamsDomain) *rest_errors.RestErr {
	filter := bson.M{"_id": idExam}
	update := bson.M{"$set": obj}

	_, err := mongodb.MongoConnection.Collection(collection_exams_name).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}
