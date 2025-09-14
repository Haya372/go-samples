package err

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	errAuthenticationFailed = errors.New("authentication failed")
)

type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(msg string) Error {
	return newError(errAuthenticationFailed, http.StatusUnauthorized, msg)
}
