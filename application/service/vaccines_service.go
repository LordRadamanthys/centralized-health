package service

import (
	"github.com/LordRadamanthys/centralized-health/adapter/output/s3"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/output"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VaccinesService struct {
	vaccinesRepository output.VaccinesPort
}

func NewVaccinesService(repository output.VaccinesPort) *VaccinesService {
	return &VaccinesService{
		vaccinesRepository: repository,
	}
}

func (vs *VaccinesService) GetVaccineByUser(userID string) ([]*domain.VaccinesDomain, *rest_errors.RestErr) {
	idConv, errConv := primitive.ObjectIDFromHex(userID)

	if errConv != nil {
		return nil, rest_errors.NewBadRequestError("invalid id")
	}

	return vs.vaccinesRepository.GetVaccineByUser(idConv)
}

func (vs *VaccinesService) GetVaccineByID(vaccineID string) (*domain.VaccinesDomain, *rest_errors.RestErr) {
	idConv, errConv := primitive.ObjectIDFromHex(vaccineID)

	if errConv != nil {
		return nil, rest_errors.NewBadRequestError("invalid id")
	}

	return vs.vaccinesRepository.GetVaccineById(idConv)
}

func (vs *VaccinesService) CreateVaccine(userID string, obj *domain.VaccinesDomain) *rest_errors.RestErr {
	idConv, errConv := primitive.ObjectIDFromHex(userID)

	if errConv != nil {
		return rest_errors.NewBadRequestError("invalid id")
	}

	obj.IdUser = idConv
	return vs.vaccinesRepository.CreateVaccine(obj)
}

func (vs *VaccinesService) InsertVaccineDocument(userID string, idVaccine string, fileName string, bufferFile []byte) *rest_errors.RestErr {
	idVaccineConv, errVaccineConv := primitive.ObjectIDFromHex(idVaccine)

	if errVaccineConv != nil {
		return rest_errors.NewBadRequestError("invalid vaccine id")
	}

	vaccine, errFind := vs.vaccinesRepository.GetVaccineById(idVaccineConv)
	if errFind != nil {
		return errFind
	}

	url, err := s3.UploadObject(userID, "vaccines", fileName, bufferFile)
	if err != nil {
		return rest_errors.NewBadRequestError("error to upload file")
	}

	vaccine.Documents = append(vaccine.Documents, domain.DocumentsDomain{Title: fileName, URLDocument: url})

	if err := vs.vaccinesRepository.UpdateVaccine(idVaccineConv, vaccine); err != nil {
		return err
	}
	return nil
}
