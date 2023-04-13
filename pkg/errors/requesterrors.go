package errors

type BadRequestError struct {
	s string
}

func (e BadRequestError) Error() string {
	return e.s
}

func NewBadRequestError(s string) error {
	return BadRequestError{s}
}
