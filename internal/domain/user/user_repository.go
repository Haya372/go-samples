package user

import (
	"context"

	"github.com/Haya372/go-samples/internal/domain/vo"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, userId vo.UserId) (*User, error)
}
