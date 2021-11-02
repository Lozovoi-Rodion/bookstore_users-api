package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		message,
		http.StatusBadRequest,
		"bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		message,
		http.StatusNotFound,
		"not_request",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		message,
		http.StatusInternalServerError,
		"internal_server_error",
	}
}
