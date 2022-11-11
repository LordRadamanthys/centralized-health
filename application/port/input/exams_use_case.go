package input

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type ExamsUseCase interface {
	GetExamsByUserID(string) ([]*domain.ExamsDomain, *rest_errors.RestErr)
	GetExamByID(string) (*domain.ExamsDomain, *rest_errors.RestErr)
	CreateExam(string, *domain.ExamsDomain) *rest_errors.RestErr
	InsertExamDocument(string, string, string, []byte) *rest_errors.RestErr
}
