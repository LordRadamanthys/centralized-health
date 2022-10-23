package repository

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type examsRepository struct{}

func NewExamsRepository() *examsRepository {
	return &examsRepository{}
}

func (*examsRepository) GetExamByUserID(string) (*domain.ExamsDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (*examsRepository) CreateExam(*domain.ExamsDomain) (*domain.ExamsDomain, *rest_errors.RestErr) {
	return nil, nil
}
