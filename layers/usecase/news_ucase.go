package usecase

import (
	"microservice/app/core"
	"microservice/layers/domain"

	"github.com/pkg/errors"
)

type NewsUseCase struct {
	log  core.Logger
	repo domain.NewsRepository
}

func NewNewsUseCase(log core.Logger, repo domain.NewsRepository) *NewsUseCase {
	return &NewsUseCase{
		log:  log,
		repo: repo,
	}
}

func (ucase *NewsUseCase) GetNews(page int32) (*domain.GetNewsResponse, error) {
	news, err := ucase.repo.FetchByPageNumber(page)
	if err != nil {
		return &domain.GetNewsResponse{}, errors.Wrap(err, "Info")
	}

	return &domain.GetNewsResponse{
		News: news,
	}, nil
}
