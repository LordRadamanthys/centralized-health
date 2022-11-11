package output

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExamsPort interface {
	GetExamsByUserID(id primitive.ObjectID) ([]*domain.ExamsDomain, *rest_errors.RestErr)
	GetExamByID(id primitive.ObjectID) (*domain.ExamsDomain, *rest_errors.RestErr)
	CreateExam(*domain.ExamsDomain) *rest_errors.RestErr
	UpdateExam(idExam primitive.ObjectID, obj *domain.ExamsDomain) *rest_errors.RestErr
}
