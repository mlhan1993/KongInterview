package utils

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/mlhan1993/KongInterview/pkg/errors"
)

// ResponseFromError writes appropriate http response for different types of errors
func ResponseFromError(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	switch err.(type) {
	case errors.BadRequestError:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
		err = fmt.Errorf("internal server error")
	}
	http.Error(w, err.Error(), statusCode)
}

func LogRequestError(err error, r *http.Request) {
	requestID := r.Context().Value("request_id")
	logger := log.WithField("request_id", requestID)
	logger.Error(err.Error())
}
