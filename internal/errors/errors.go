package errors

import (
	"log"
)

type ErrorResponse struct {
	Code    string
	Status  int
	Message string
	Details interface{}
}

const (
	ErrCodeBadRequest    = "BAD_REQUEST"
	ErrCodeForbidden     = "FORBIDDEN"
	ErrCodeInternalError = "INTERNAL_ERROR"
	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeUnauthorized  = "UNAUTHORIZED"
)

func NewErrorResponse(status int, code, message string, details ...interface{}) *ErrorResponse {
	ae := &ErrorResponse{
		Status:  status,
		Code:    code,
		Message: message,
	}

	if len(details) > 0 {
		ae.Details = details[0]
	}
	return ae
}

// Common Error
var (
	ErrNotFound        = NewErrorResponse(404, ErrCodeNotFound, "Resource not found")
	ErrBadRequest      = NewErrorResponse(400, ErrCodeBadRequest, "Bad request")
	ErrForbidden       = NewErrorResponse(403, ErrCodeForbidden, "Forbidden access")
	ErrUnauthorized    = NewErrorResponse(401, ErrCodeUnauthorized, "Unauthorized access")
	ErrInternalRequest = NewErrorResponse(500, ErrCodeInternalError, "An unexpected error happened")
)

func ErrorHandler(err error) *ErrorResponse {
	log.Println(err)
	switch e := err.(type) {
	case *AlreadyExistError:
		response := ErrBadRequest
		response.Message = e.Message
		return response
	case *NotFoundError:
		response := ErrNotFound
		response.Message = e.Message
		return response
	case FieldErrors:
		return NewValidationErrorResponse(e)
	default:
		return ErrInternalRequest
	}
}
