package di

import "go.uber.org/dig"

func NewDiContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(newEngine); err != nil {
		return nil, err
	}

	if err := provideControllers(container); err != nil {
		return nil, err
	}

	if err := provideMiddlewares(container); err != nil {
		return nil, err
	}

	if err := provideRepositories(container); err != nil {
		return nil, err
	}

	if err := provideServices(container); err != nil {
		return nil, err
	}

	if err := provideUsecases(container); err != nil {
		return nil, err
	}

	return container, nil
}
