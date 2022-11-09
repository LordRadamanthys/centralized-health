package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func RoutesUrl() {
	dependecies := LoadDependecies()

	router.GET("/ping", ping)

	router.POST("/user", dependecies.userController.CreateUser)
	router.PUT("/user", dependecies.userController.UpdateUserByID)
	router.POST("/login", dependecies.userController.LoginController)

	router.Run(":8081")
}

func ping(ctx *gin.Context) {
	ctx.Writer.WriteString("Pong")
}
