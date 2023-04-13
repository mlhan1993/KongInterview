package utils

import (
	"github.com/mlhan1993/KongInterview/pkg/errors"
	"net/http"
)

func ResponseFromError(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	switch err.(type) {
	case errors.BadRequestError:
		statusCode = http.StatusBadRequest
	}
	http.Error(w, err.Error(), statusCode)
}
