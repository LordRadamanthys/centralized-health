package middleware

import (
	"strings"

	jwtconfig "github.com/LordRadamanthys/centralized-health/configuration/jwt_config"
	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"github.com/gin-gonic/gin"
)

type auth struct{}

type authInterface interface {
	Auth() gin.HandlerFunc
}

var Auth authInterface = &auth{}

func (a *auth) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Beare_schema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(401, rest_errors.NewUnauthorizedError("invalid token"))
		}

		token := strings.Split(header, Beare_schema)[1]
		if !jwtconfig.NewJWTUtils().ValidateToken(token) {
			c.AbortWithStatusJSON(401, rest_errors.NewUnauthorizedError("invalid token"))
		}
	}
}
