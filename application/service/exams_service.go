package service

import (
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/output"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExamsService struct {
	examsRepository output.ExamsPort
}

func NewExamsService(repository output.ExamsPort) *ExamsService {
	return &ExamsService{
		examsRepository: repository,
	}
}

func (es *ExamsService) GetExamByUserID(id string) (*domain.ExamsDomain, *rest_errors.RestErr) {
	idConv, errConv := primitive.ObjectIDFromHex(id)

	if errConv != nil {
		return nil, rest_errors.NewBadRequestError("invalid id")
	}

	return es.examsRepository.GetExamByUserID(idConv)
}

func (es *ExamsService) CreateExam(*domain.ExamsDomain) (*domain.ExamsDomain, *rest_errors.RestErr) {
	return nil, nil
}
