package usecase

import (
	"microservice/app/core"
	"microservice/layers/domain"
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

func (ucase *NewsUseCase) GetNews(page int32) ([]*domain.NewsCard, error) {

	panic("NOT IMPLEMENTS")
	//news, err := ucase.repo.FetchByPageNumber(page)
	//if err != nil {
	//	return &domain.GetNewsResponse{}, errors.Wrap(err, "Info")
	//}
	//
	//return &domain.GetNewsResponse{
	//	News: news,
	//}, nil
}
