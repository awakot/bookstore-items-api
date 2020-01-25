package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad reques",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedError() *RestErr {
	return &RestErr{
		Message: "unable to retrieve user information from given access_token",
		Code:    http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
