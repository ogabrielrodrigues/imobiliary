package httperr

import (
	"time"
)

const (
	BadRequestError          string = "BAD_REQUEST_ERROR"
	UnauthorizedError        string = "UNAUTHORIZED_ERROR"
	ForbiddenError           string = "FORBIDDEN_ERROR"
	NotFoundError            string = "NOT_FOUND_ERROR"
	AlreadyExistsError       string = "ALREADY_EXISTS_ERROR"
	UnprocessableEntityError string = "UNPROCESSABLE_ENTITY_ERROR"
	InternalServerError      string = "INTERNAL_SERVER_ERROR"
)

type HttpError struct {
	Code      string
	Message   string
	Context   map[string]interface{}
	Timestamp time.Time
	Cause     error
}

func (err HttpError) Error() string {
	return err.Message
}

func (err HttpError) Unwrap() error {
	return err.Cause
}

func (e *HttpError) Is(target error) bool {
	if targetErr, ok := target.(*HttpError); ok {
		return e.Code == targetErr.Code
	}

	return false
}
