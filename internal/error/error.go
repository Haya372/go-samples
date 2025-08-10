package error

type Error interface {
	Error() string
	Status() int
	Message() string
}

type errorImpl struct {
	status  int
	message string
	error
}

func (e *errorImpl) Status() int {
	return e.status
}

func (e *errorImpl) Message() string {
	return e.message
}

func newError(err error, status int, message string) Error {
	return &errorImpl{
		error:   err,
		status:  status,
		message: message,
	}
}
