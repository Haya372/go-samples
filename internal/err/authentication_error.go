package err

import (
	"net/http"
)

const (
	AuthenticationFailed ErrorCode = "AUTHENTICATION_FAILED"
)

type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(msg string) Error {
	return newError(AuthenticationFailed, http.StatusUnauthorized, msg)
}
