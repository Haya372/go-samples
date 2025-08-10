package di

import (
	"github.com/Haya372/go-samples/internal/gateway/memory/repository/user"
	"go.uber.org/dig"
)

func provideRepositories(c *dig.Container) error {
	if err := c.Provide(user.NewUserRepository); err != nil {
		return err
	}
	return nil
}
