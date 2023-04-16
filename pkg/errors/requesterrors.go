package errors

import "encoding/json"

type ErrorMessage struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"statusCode"`
	// Code int `json:"code"` # potentially can be introduced along with more special errors
}

func (e ErrorMessage) ToString() string {
	byteMsg, _ := json.Marshal(e)
	return string(byteMsg)
}

type BadRequestError struct {
	s string
}

func (e BadRequestError) Error() string {
	return e.s
}

func NewBadRequestError(s string) error {
	return BadRequestError{s}
}
