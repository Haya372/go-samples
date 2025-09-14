package di

import (
	"time"

	"github.com/Haya372/go-samples/internal/controller/middleware"
	"go.uber.org/dig"
)

const defaultRefreshTimeoutHour = 24

func provideMiddlewares(container *dig.Container) error {
	if err := container.Provide(middleware.NewGinJwtMiddleware); err != nil {
		return err
	}

	// TODO: remove this provides after implement config loader
	if err := container.Provide(func() middleware.JwtConfig {
		return middleware.JwtConfig{
			Realm:          "test",
			SecretKey:      "test secret",
			Timeout:        1 * time.Hour,
			RefreshTimeout: defaultRefreshTimeoutHour * time.Hour,
		}
	}); err != nil {
		return err
	}

	return nil
}
