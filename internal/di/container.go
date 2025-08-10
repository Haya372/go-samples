package di

import "go.uber.org/dig"

func NewDiContainer() (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(NewEngine); err != nil {
		return nil, err
	}

	if err := ProvideControllers(c); err != nil {
		return nil, err
	}

	if err := provideMiddlewares(c); err != nil {
		return nil, err
	}

	if err := provideRepositories(c); err != nil {
		return nil, err
	}

	if err := provideServices(c); err != nil {
		return nil, err
	}

	if err := provideUsecases(c); err != nil {
		return nil, err
	}

	return c, nil
}
