package input

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type ExamsUseCase interface {
	GetExamByUser(string) (*domain.ExamsDomain, *rest_errors.RestErr)
	CreateExam(string, *domain.ExamsDomain) *rest_errors.RestErr
}
