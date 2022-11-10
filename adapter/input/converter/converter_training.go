package converter

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/jinzhu/copier"
)

func ConverterRequestIntoDomainTraining(request *requests.TrainingRequest) *domain.TrainingDomain {

	var domain domain.TrainingDomain

	copier.Copy(&domain, request)

	return &domain
}
