package di

import (
	"time"

	"github.com/Haya372/go-samples/internal/controller/middleware"
	"go.uber.org/dig"
)

func provideMiddlewares(c *dig.Container) error {
	if err := c.Provide(middleware.NewGinJwtMiddleware); err != nil {
		return err
	}

	// TODO: remove this provides after implement config loader
	if err := c.Provide(func() middleware.JwtConfig {
		return middleware.JwtConfig{
			Realm:          "test",
			SecretKey:      "test secret",
			Timeout:        1 * time.Hour,
			RefreshTimeout: 24 * time.Hour,
		}
	}); err != nil {
		return err
	}

	return nil
}
