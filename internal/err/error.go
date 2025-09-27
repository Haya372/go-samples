package err

import "fmt"

type ErrorCode string

type Error interface {
	Error() string
	Code() ErrorCode
	Status() int
	Message() string
}

type customError struct {
	code    ErrorCode
	status  int
	message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%v: %s", e.code, e.message)
}

func (e *customError) Code() ErrorCode {
	return e.code
}

func (e *customError) Status() int {
	return e.status
}

func (e *customError) Message() string {
	return e.message
}

func newError(code ErrorCode, status int, message string) Error {
	return &customError{
		code:    code,
		status:  status,
		message: message,
	}
}
