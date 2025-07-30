package errors

type FieldErrors map[string]string

func (e FieldErrors) Error() string {
	return "test"
}

func NewValidationErrorResponse(fe FieldErrors) *ErrorResponse {
	return NewErrorResponse(400, "INVALID_INPUT", "Validation failed", fe)
}
