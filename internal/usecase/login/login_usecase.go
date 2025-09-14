package login

import (
	"context"
	"fmt"

	"github.com/Haya372/go-samples/internal/domain/service"
	"github.com/Haya372/go-samples/internal/domain/user"
	"github.com/Haya372/go-samples/internal/domain/vo"

	customErr "github.com/Haya372/go-samples/internal/err"
)

type LoginParam struct {
	UserId      vo.UserId
	RawPassword string
}

type LoginUsecase interface {
	Execute(ctx context.Context, param LoginParam) (*user.User, error)
}

type loginUsecaseImpl struct {
	userRepository          user.UserRepository
	userAuthenticateService service.UserAuthenticateService
}

func (uc *loginUsecaseImpl) Execute(ctx context.Context, param LoginParam) (*user.User, error) {
	user, err := uc.userRepository.FindByEmail(ctx, param.UserId)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, customErr.NewAuthenticationError(fmt.Sprintf("user not found, userId=%v", param.UserId))
	}

	if ok := uc.userAuthenticateService.Execute(ctx, *user, param.RawPassword); !ok {
		return nil, customErr.NewAuthenticationError(fmt.Sprintf("invalid password, userId=%v", param.UserId))
	}

	return user, nil
}

func NewLoginUsecase(
	userRepository user.UserRepository,
	userAuthenticateService service.UserAuthenticateService) LoginUsecase {
	return &loginUsecaseImpl{
		userRepository:          userRepository,
		userAuthenticateService: userAuthenticateService,
	}
}
