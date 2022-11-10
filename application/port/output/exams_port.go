package output

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExamsPort interface {
	GetExamByUserID(primitive.ObjectID) (*domain.ExamsDomain, *rest_errors.RestErr)
	CreateExam(*domain.ExamsDomain) (*domain.ExamsDomain, *rest_errors.RestErr)
}
