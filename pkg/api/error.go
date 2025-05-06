package api

import (
	"net/http"
	"restfullApi/pkg/errors"
)

func result(err *errors.Error) any {
	if err == nil {

	}
	return nil
}

func err2code(err *errors.Error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err.ErrorType {
	case errors.Token:
		return http.StatusUnauthorized
	case errors.User, errors.Validation:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}

}
