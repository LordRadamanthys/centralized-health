package service

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/domain"
	"github.com/LordRadamanthys/centralized-health/application/port/output"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"github.com/LordRadamanthys/centralized-health/configuration/utils/criptography"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	userRepository output.UserPort
}

func NewUserService(repository output.UserPort) *userService {
	return &userService{
		userRepository: repository,
	}
}

func (us *userService) GetUserByID(id string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (us *userService) GetUserByEmail(email string) (*domain.UserDomain, *rest_errors.RestErr) {
	return nil, nil
}

func (us *userService) UpdateUserByID(id primitive.ObjectID, user *requests.UserRequest) *rest_errors.RestErr {
	return nil
}

func (us *userService) UpdateUserByEmail(user *requests.UserRequest) *rest_errors.RestErr {
	return nil
}

func (us *userService) CreateUser(user *requests.UserRequest) *rest_errors.RestErr {
	user.Password = criptography.EncodePassword(user.Password)
	return us.userRepository.CreateUser(user)
}

func (us *userService) LoginService(email string, password string) (*domain.UserDomain, *rest_errors.RestErr) {
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !criptography.ComparePassword(password, []byte(user.Password)) {
		return nil, rest_errors.NewBadRequestError("invalid user")
	}
	return user, nil
}
