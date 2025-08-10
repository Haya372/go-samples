package error

import "net/http"

type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(err error) Error {
	return newError(err, http.StatusUnauthorized, "illegal userId or password")
}
