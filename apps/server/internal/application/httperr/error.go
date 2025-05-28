package httperr

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	appcontext "imobiliary/internal/application/context"
)

const (
	ValidationError          string = "VALIDATION_ERROR"
	BadRequestError          string = "BAD_REQUEST_ERROR"
	UnauthorizedError        string = "UNAUTHORIZED_ERROR"
	ForbiddenError           string = "FORBIDDEN_ERROR"
	NotFoundError            string = "NOT_FOUND_ERROR"
	AlreadyExistsError       string = "ALREADY_EXISTS_ERROR"
	UnprocessableEntityError string = "UNPROCESSABLE_ENTITY_ERROR"
	InternalServerError      string = "INTERNAL_SERVER_ERROR"
)

type HttpError struct {
	Code      string                 `json:"code"`
	HttpCode  int                    `json:"http_code"`
	Message   string                 `json:"message"`
	Context   map[string]interface{} `json:"context,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
	Cause     error                  `json:"cause"`
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

func NewValidationError(ctx context.Context, validationErrors *ValidationErrors) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:      ValidationError,
		HttpCode:  http.StatusUnprocessableEntity,
		Message:   "One or more validation errors occurred",
		Context:   map[string]interface{}{},
		Timestamp: currentTime,
		Cause:     validationErrors,
	}
}

func NewInternalServerError(ctx context.Context, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     InternalServerError,
		HttpCode: http.StatusInternalServerError,
		Message:  "Internal server error",
		Context: map[string]interface{}{
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func NewInvalidContextError(ctx context.Context, fieldName string, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     InternalServerError,
		HttpCode: http.StatusInternalServerError,
		Message:  fmt.Sprintf("Valor de '%s' inválido: %s", fieldName, reason),
		Context: map[string]interface{}{
			"field":  fieldName,
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func NewBadRequestError(ctx context.Context, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     BadRequestError,
		HttpCode: http.StatusBadRequest,
		Message:  "Bad request",
		Context: map[string]interface{}{
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func NewUnauthorizedError(ctx context.Context, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     UnauthorizedError,
		HttpCode: http.StatusUnauthorized,
		Message:  "Unauthorized",
		Context: map[string]interface{}{
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func NewForbiddenError(ctx context.Context, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     ForbiddenError,
		HttpCode: http.StatusForbidden,
		Message:  "Forbidden",
		Context: map[string]interface{}{
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func NewNotFoundError(ctx context.Context, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     NotFoundError,
		HttpCode: http.StatusNotFound,
		Message:  "Not found",
		Context: map[string]interface{}{
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func NewAlreadyExistsError(ctx context.Context, reason string) *HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		currentTime = time.Time{}
	}

	return &HttpError{
		Code:     AlreadyExistsError,
		HttpCode: http.StatusConflict,
		Message:  "Conflict",
		Context: map[string]interface{}{
			"reason": reason,
		},
		Timestamp: currentTime,
	}
}

func GetValidationErrors(err error) *ValidationErrors {
	var appErr *HttpError
	if errors.As(err, &appErr) && appErr.Code == ValidationError {
		var validationErrs *ValidationErrors
		if errors.As(appErr.Cause, &validationErrs) {
			return validationErrs
		}
	}
	return nil
}

func GetErrorContext(err error) (map[string]interface{}, bool) {
	var appErr *HttpError
	if errors.As(err, &appErr) {
		return appErr.Context, true
	}
	return nil, false
}
