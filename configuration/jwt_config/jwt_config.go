package jwtconfig

import (
	"fmt"
	"strings"
	"time"

	"github.com/LordRadamanthys/centralized-health/configuration/rest_errors"
	"github.com/dgrijalva/jwt-go"
)

const Beare_schema = "Bearer "

type jwtObject struct {
	secretKey string
	issure    string
}

func NewJWTUtils() *jwtObject {
	return &jwtObject{
		secretKey: "secret-key",
		issure:    "centrilized_health-ms",
	}

}

type Claim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func (s *jwtObject) GeneratedToken(id string) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}

	return t, nil
}

func (s *jwtObject) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *jwtObject) GetId(tokenJWT string) (string, *rest_errors.RestErr) {
	if tokenJWT == "" {
		return "", rest_errors.NewBadRequestError("authorization not found")
	}
	token := strings.Split(strings.TrimSpace(tokenJWT), Beare_schema)[1]
	claims := &Claim{}

	_, err := jwt.ParseWithClaims(strings.TrimSpace(token), claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(strings.TrimSpace(s.secretKey)), nil
	})

	if err != nil {
		return "", rest_errors.NewInternalServerError("error to get id", nil)
	}

	return claims.Id, nil
}
