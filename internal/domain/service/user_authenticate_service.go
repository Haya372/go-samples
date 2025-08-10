package service

import (
	"context"

	"github.com/Haya372/go-samples/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthenticateService interface {
	Execute(ctx context.Context, user user.User, rawPassword string) bool
}

type userAuthenticateServiceImpl struct {
}

func (s *userAuthenticateServiceImpl) Execute(ctx context.Context, user user.User, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword))
	return err == nil
}

func NewUserAuthenticateService() UserAuthenticateService {
	return &userAuthenticateServiceImpl{}
}
