package errors

import "fmt"

type AlreadyExistError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("[not found error]: %v", e.Message)
	}
	return "[resource not found]"
}

func (e *AlreadyExistError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("[already exist error]: , %v", e.Message)
	}
	return "[already exist]"
}
