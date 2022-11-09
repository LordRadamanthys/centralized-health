package service

import (
	"fmt"

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
	user, err := us.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) GetUserByEmail(email string) (*domain.UserDomain, *rest_errors.RestErr) {
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) UpdateUserByID(id primitive.ObjectID, user *requests.UserRequest) *rest_errors.RestErr {
	userFind, err := us.GetUserByID(id.Hex())

	if err != nil {
		return err
	}

	if userFind == nil {
		return rest_errors.NewBadRequestError("user not found")
	}

	userFind.Birth = user.Birth
	// userFind.Email = user.Email
	userFind.Name = user.Name
	userFind.Phones = user.Phones

	return us.userRepository.UpdateUserByID(id, userFind)
}

func (us *userService) UpdateUserByEmail(user *requests.UserRequest) *rest_errors.RestErr {
	return nil
}

func (us *userService) CreateUser(user *requests.UserRequest) *rest_errors.RestErr {

	findUser, _ := us.userRepository.GetUserByEmail(user.Email)
	fmt.Println(findUser)
	if findUser != nil {
		return rest_errors.NewBadRequestError("user already exist")
	}

	user.Password = criptography.EncodePassword(user.Password)
	us.userRepository.CreateUser(user)

	return nil
}

func (us *userService) LoginService(email string, password string) (*domain.UserDomain, *rest_errors.RestErr) {
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !criptography.ComparePassword(password, []byte(user.Password)) {
		return nil, rest_errors.NewNotFoundError("user not found")
	}
	return user, nil
}
