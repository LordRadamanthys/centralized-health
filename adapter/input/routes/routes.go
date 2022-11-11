package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func RoutesUrl() {
	dependecies := LoadDependecies()

	router.GET("/ping", ping)

	router.POST("/user", dependecies.userController.CreateUser)
	router.PUT("/user", dependecies.userController.UpdateUserByID)
	router.POST("/login", dependecies.userController.LoginController)

	router.POST("/vaccines", dependecies.vaccinesController.CreateVaccine)
	router.GET("/vaccines", dependecies.vaccinesController.GetVaccineByUser)

	router.GET("/training", dependecies.trainingController.GetTraining)
	router.POST("/training", dependecies.trainingController.CreateTraining)
	router.PUT("/training/:day", dependecies.trainingController.UpdateTraining)

	router.GET("/exams", dependecies.examsController.GetExamsByUser)
	router.POST("/exams", dependecies.examsController.CreateExam)

	router.Run(":8081")
}

func ping(ctx *gin.Context) {
	ctx.Writer.WriteString("Pong")
}
