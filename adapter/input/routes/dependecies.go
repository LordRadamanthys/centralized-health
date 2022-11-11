package routes

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/controller"
	"github.com/LordRadamanthys/centralized-health/adapter/output/repository"
	"github.com/LordRadamanthys/centralized-health/application/service"
)

type DependeciesRoutes struct {
	userController     *controller.UserController
	examsController    *controller.ExamsController
	vaccinesController *controller.VaccinesController
	trainingController *controller.TrainingController
}

func LoadDependecies() *DependeciesRoutes {

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	examsRepository := repository.NewExamsRepository()
	examsService := service.NewExamsService(examsRepository)
	examsController := controller.NewExamsService(examsService)

	trainingRepository := repository.NewTrainingRepository()
	trainingService := service.NewTrainingService(trainingRepository)
	trainingController := controller.NewTrainingService(trainingService)

	vaccinesRepository := repository.NewVaccinesRepository()
	vaccinesService := service.NewVaccinesService(vaccinesRepository)
	vaccinesController := controller.NewVaccinesController(vaccinesService)

	return &DependeciesRoutes{
		userController:     userController,
		examsController:    examsController,
		vaccinesController: vaccinesController,
		trainingController: trainingController,
	}
}
