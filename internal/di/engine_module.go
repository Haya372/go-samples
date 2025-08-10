package di

import (
	"github.com/Haya372/go-samples/internal/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func provideControllers(c *dig.Container) error {
	if err := c.Provide(controller.NewHealthCheckController, dig.Group("controller")); err != nil {
		return err
	}

	if err := c.Provide(controller.NewLoginController, dig.Group("controller")); err != nil {
		return err
	}
	return nil
}

type controllerParams struct {
	dig.In

	Controllers []controller.Controller `group:"controller"`
}

func newEngine(p controllerParams) *gin.Engine {
	r := gin.Default()
	for _, c := range p.Controllers {
		c.SetRoute(r)
	}
	return r
}
