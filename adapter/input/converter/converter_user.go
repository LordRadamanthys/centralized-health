package converter

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/adapter/input/response"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/jinzhu/copier"
)

func ConverterDomainIntoRequestUser(domain *domain.UserDomain) *requests.UserRequest {

	var request requests.UserRequest

	copier.Copy(&request, domain)
	return &request

}

func ConverterRequestIntoDomainUser(request *requests.UserRequest) *domain.UserDomain {

	var domain domain.UserDomain

	copier.Copy(&domain, request)

	return &domain
}

func ConverterDomainIntoResponseUser(domain *domain.UserDomain) *response.UserResponse {

	var response response.UserResponse

	copier.Copy(&response, domain)

	return &response
}
