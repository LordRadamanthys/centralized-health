package repository

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type examsRepository struct{}

func NewExamsRepository() *examsRepository {
	return &examsRepository{}
}

func (*examsRepository) GetExamByUserID(id primitive.ObjectID) (*domain.ExamsDomain, *rest_errors.RestErr) {

	return nil, nil
}

func (*examsRepository) CreateExam(*domain.ExamsDomain) *rest_errors.RestErr {
	return nil
}
