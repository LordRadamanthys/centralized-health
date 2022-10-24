package routes

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/controller"
	"github.com/LordRadamanthys/centralized-health/adapter/output/repository"
	"github.com/LordRadamanthys/centralized-health/application/service"
)

type DependeciesRoutes struct {
	userController *controller.UserController
}

func LoadDependecies() *DependeciesRoutes {

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	return &DependeciesRoutes{
		userController: userController,
	}
}
