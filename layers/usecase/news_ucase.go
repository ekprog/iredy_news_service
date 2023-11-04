package usecase

import (
	"microservice/app/core"
	"microservice/layers/domain"
	pb "microservice/pkg/pb/api"

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

func (ucase *NewsUseCase) GetNews(page int32) (domain.GetNewsResponse, error) {
	// Если передали страницу <= 0, не выходим из функции
	if page <= 0 {
		return domain.GetNewsResponse{
			Status: pb.Status{
				Code:    domain.ValidationError,
				Message: "page can't have value of <= 0",
			},
			News: []*domain.NewsCard{},
		}, nil
	}

	repoRes, err := ucase.repo.FetchByPageNumber(page)
	// Ошибка запроса к базе
	if err != nil {
		return domain.GetNewsResponse{}, errors.Wrap(err, "FetchByPageNumber")
	}

	// Null ответ от базы

	// ToDo: Здесь нет err (вместо errors.Wrap должен быть nil)
	// (Либо возвращаем ошибку, либо структуру - и то и тл нет смысла возвращать, потому что если есть
	// ошибка, то все остальное игнорируется)
	if repoRes == nil {
		return domain.GetNewsResponse{
			Status: pb.Status{
				Code:    domain.NotFound,
				Message: "There are no news",
			},
			News: []*domain.NewsCard{},
		}, errors.Wrap(err, "FetchByPageNumber")
	}

	// Успех
	return domain.GetNewsResponse{
		Status: pb.Status{
			Code:    domain.Success,
			Message: domain.Success,
		},
		News: repoRes,
	}, nil

}

func (ucase *NewsUseCase) AddNewsCard(newsCard domain.NewsCard) (domain.CreateNewsResponse, error) {

	return domain.CreateNewsResponse{
		Status: pb.Status{
			Code:    domain.Success,
			Message: domain.Success,
		},
		Id: newsCard.Id,
	}, nil
}
