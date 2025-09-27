package user

import (
	"context"

	"github.com/Haya372/go-samples/internal/domain/user"
	"github.com/Haya372/go-samples/internal/domain/vo"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryImpl struct{}

func (repo *userRepositoryImpl) FindByEmail(ctx context.Context, userId vo.UserId) (*user.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	users := map[string]user.User{
		"user-id-1": {
			Id:       vo.NewUserId("user-id-1"),
			Name:     "Test User",
			Password: vo.NewPassword(string(hash)),
		},
	}

	u, exists := users[userId.Value()]
	if !exists {
		return nil, nil
	}

	return &u, nil
}

func NewUserRepository() user.UserRepository {
	return &userRepositoryImpl{}
}
