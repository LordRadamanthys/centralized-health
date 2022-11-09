package controller

import (
	"fmt"
	"net/http"

	"github.com/LordRadamanthys/centralized-health/adapter/input/converter"
	"github.com/LordRadamanthys/centralized-health/adapter/input/requests"
	"github.com/LordRadamanthys/centralized-health/application/port/input"
	jwtconfig "github.com/LordRadamanthys/centralized-health/configuration/jwt_config"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userService input.UserUseCase
}

func NewUserController(service input.UserUseCase) *UserController {
	return &UserController{
		userService: service,
	}
}

func (uc *UserController) LoginController(ctx *gin.Context) {
	var loginObj *requests.LoginRequest

	if err := ctx.ShouldBindJSON(&loginObj); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := uc.userService.LoginService(loginObj.Email, loginObj.Password)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	fmt.Println(user)
	token, errJwt := jwtconfig.NewJWTUtils().GeneratedToken(user.ID.Hex())

	if errJwt != nil {
		ctx.JSON(400, "error to generate token")
		return
	}

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	userConv := converter.ConverterDomainIntoResponseUser(user)
	ctx.Writer.Header().Add("token", token)
	ctx.JSON(http.StatusOK, userConv)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user *requests.UserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, "invalid object")
	}

	if err := uc.userService.CreateUser(user); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("invalid id"))
	}

	user, err := uc.userService.GetUserByID(id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(200, user)
}

func (uc *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("invalid email"))
	}

	user, err := uc.userService.GetUserByEmail(email)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(200, user)
}

func (uc *UserController) UpdateUserByID(ctx *gin.Context) {
	var user *requests.UserRequest
	header := ctx.GetHeader("Authorization")
	id, err := jwtconfig.NewJWTUtils().GetId(header)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, "invalid object")
		return
	}

	fmt.Println(id)
	idConv, errConv := primitive.ObjectIDFromHex(string(id))

	if errConv != nil {
		ctx.JSON(400, rest_errors.NewInternalServerError("invalid id: "+errConv.Error(), nil))
		return
	}

	errUpdate := uc.userService.UpdateUserByID(idConv, user)
	if errUpdate != nil {
		ctx.JSON(errUpdate.Code, errUpdate)
		return
	}
	ctx.JSON(200, nil)
}
