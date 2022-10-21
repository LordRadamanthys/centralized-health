package rest_errors

import (
	"errors"
	"net/http"
)

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RestErr struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Error   string  `json:"error"`
	Causes  []Cause `json:"causes"`
}

func NewRestError(message string, code int, err string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Code:    code,
		Error:   err,
		Causes:  causes,
	}
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestValidationError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	return result
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}
