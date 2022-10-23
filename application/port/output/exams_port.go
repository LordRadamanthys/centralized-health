package output

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
)

type ExamsPort interface {
	GetExamByUserID(string) (*domain.ExamsDomain, *rest_errors.RestErr)
	CreateExam(*domain.ExamsDomain) (*domain.ExamsDomain, *rest_errors.RestErr)
}
