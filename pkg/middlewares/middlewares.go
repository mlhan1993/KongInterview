package middlewares

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const RequestIDKey = "request_id"

// RequestIDMiddleware wrap http request with a unique id
func RequestIDMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, RequestIDKey, id.String())
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

// LogRequestMiddleware log the request whenever we receive one
func LogRequestMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(RequestIDKey)
		rLogger := log.WithFields(log.Fields{
			"request_id": requestID,
			"uri":        r.RequestURI,
			"header":     r.Header,
			"method":     r.Method,
		})

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			rLogger.Error(err.Error())
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		rLogger.WithField("body", string(body))
		rLogger.Info("received")

		handler.ServeHTTP(w, r)
	})
}
