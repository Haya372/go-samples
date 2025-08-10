package di

import (
	"github.com/Haya372/go-samples/internal/usecase/login"
	"go.uber.org/dig"
)

func provideUsecases(c *dig.Container) error {
	if err := c.Provide(login.NewLoginUsecase); err != nil {
		return err
	}
	return nil
}
