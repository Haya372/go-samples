package di

import (
	"github.com/Haya372/go-samples/internal/domain/service"
	"go.uber.org/dig"
)

func provideServices(c *dig.Container) error {
	if err := c.Provide(service.NewUserAuthenticateService); err != nil {
		return err
	}

	return nil
}
