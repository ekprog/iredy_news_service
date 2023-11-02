package bootstrap

import (
	"microservice/app"
	"microservice/layers/delivery/grpc"
	"microservice/layers/domain"
	"microservice/layers/repos"
	"microservice/layers/usecase"

	"go.uber.org/dig"
)

func initDependencies(di *dig.Container) error {

	// Repository
	_ = di.Provide(repos.NewNewsrepo, dig.As(new(domain.NewsRepository)))

	// Services

	// Use Cases
	_ = di.Provide(usecase.NewNewsUseCase, dig.As(new(domain.NewsUseCase)))

	//delivery
	if err := app.InitDelivery(grpc.NewNewsService); err != nil {
		return err
	}
	return nil
}
