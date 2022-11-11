package service

import (
	"github.com/LordRadamanthys/centralized-health/adapter/output/s3"
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

func (es *ExamsService) GetExamsByUserID(id string) ([]*domain.ExamsDomain, *rest_errors.RestErr) {
	idConv, errConv := primitive.ObjectIDFromHex(id)

	if errConv != nil {
		return nil, rest_errors.NewBadRequestError("invalid id")
	}

	return es.examsRepository.GetExamsByUserID(idConv)
}

func (es *ExamsService) GetExamByID(id string) (*domain.ExamsDomain, *rest_errors.RestErr) {
	idConv, errConv := primitive.ObjectIDFromHex(id)

	if errConv != nil {
		return nil, rest_errors.NewBadRequestError("invalid id")
	}

	return es.examsRepository.GetExamByID(idConv)
}

func (es *ExamsService) CreateExam(id string, obj *domain.ExamsDomain) *rest_errors.RestErr {

	idConv, errConv := primitive.ObjectIDFromHex(id)

	if errConv != nil {
		return rest_errors.NewBadRequestError("invalid id")
	}
	obj.IdUser = idConv
	return es.examsRepository.CreateExam(obj)
}

func (es *ExamsService) InsertExamDocument(userID string, idExam string, fileName string, bufferFile []byte) *rest_errors.RestErr {
	idExamConv, errVaccineConv := primitive.ObjectIDFromHex(idExam)

	if errVaccineConv != nil {
		return rest_errors.NewBadRequestError("invalid vaccine id")
	}

	exam, errFind := es.examsRepository.GetExamByID(idExamConv)
	if errFind != nil {
		return errFind
	}

	url, err := s3.UploadObject(userID, "exams", fileName, bufferFile)
	if err != nil {
		return rest_errors.NewBadRequestError("error to upload file")
	}

	exam.Documents = append(exam.Documents, domain.DocumentsDomain{Title: fileName, URLDocument: url})

	if err := es.examsRepository.UpdateExam(idExamConv, exam); err != nil {
		return err
	}
	return nil
}
