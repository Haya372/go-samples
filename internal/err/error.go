package err

type Error interface {
	Error() string
	Status() int
	Message() string
}

type customError struct {
	error

	status  int
	message string
}

func (e *customError) Status() int {
	return e.status
}

func (e *customError) Message() string {
	return e.message
}

func newError(err error, status int, message string) Error {
	return &customError{
		error:   err,
		status:  status,
		message: message,
	}
}
