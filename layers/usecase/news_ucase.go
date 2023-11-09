package usecase

import (
	"context"
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

func (ucase *NewsUseCase) GetNews(ctx context.Context, page int32) (domain.GetNewsResponse, error) {
	// Если передали страницу <= 0, не выходим из функции
	if page <= 0 {
		return domain.GetNewsResponse{
			Status: domain.Status{
				Code:    domain.ValidationError,
				Message: "page can't have value of <= 0 or page is required",
			},
			News: []*domain.NewsCard{},
		}, nil
	}

	repoRes, err := ucase.repo.FetchByPageNumber(ctx, page)
	// Ошибка запроса к базе
	if err != nil {
		return domain.GetNewsResponse{}, errors.Wrap(err, "FetchByPageNumber")
	}

	// (Либо возвращаем ошибку, либо структуру - и то и тл нет смысла возвращать, потому что если есть
	// ошибка, то все остальное игнорируется)
	if repoRes == nil {
		return domain.GetNewsResponse{
			Status: domain.Status{
				Code:    domain.NotFound,
				Message: "There are no news",
			},
			News: []*domain.NewsCard{},
		}, nil
	}

	// Успех
	return domain.GetNewsResponse{
		Status: domain.Status{
			Code:    domain.Success,
			Message: domain.Success,
		},
		News: repoRes,
	}, nil

}
func (ucase *NewsUseCase) GetNewsDetails(ctx context.Context, page int32, news_id int32) (domain.GetNewsDetailsResponse, error) {
	// Если передали страницу <= 0, не выходим из функции
	if page <= 0 {
		return domain.GetNewsDetailsResponse{
			Status: domain.Status{
				Code:    domain.ValidationError,
				Message: "page can't have value of <= 0 or page is required",
			},
			NewsDetails: nil,
		}, nil
	}

	repoRes, err := ucase.repo.FetchNewsDetailsByPageNumber(ctx, page, news_id)
	// Ошибка запроса к базе
	if err != nil {
		return domain.GetNewsDetailsResponse{}, errors.Wrap(err, "FetchByPageNumber")
	}

	// (Либо возвращаем ошибку, либо структуру - и то и тл нет смысла возвращать, потому что если есть
	// ошибка, то все остальное игнорируется)
	if repoRes == nil {
		return domain.GetNewsDetailsResponse{
			Status: domain.Status{
				Code:    domain.NotFound,
				Message: "There are no newsDetails",
			},
			NewsDetails: nil,
		}, nil
	}

	//Успех
	return domain.GetNewsDetailsResponse{
		Status: domain.Status{
			Code:    domain.Success,
			Message: domain.Success,
		},
		NewsDetails: repoRes,
	}, nil

}

func (ucase *NewsUseCase) AddNewsCard(ctx context.Context, newsCard domain.NewsCard) (domain.CreateNewsResponse, error) {

	resId, err := ucase.repo.InsertIfNotExistsNewsCard(ctx, &newsCard)
	// Ошибка запроса к базе
	if err != nil {
		return domain.CreateNewsResponse{}, errors.Wrap(err, "InsertIfNotExists")
	}

	if resId == 0 {
		return domain.CreateNewsResponse{
			Status: domain.Status{
				Code:    domain.FieldRequired,
				Message: "Field Title Required",
			},
			Id: 0,
		}, nil
	}

	return domain.CreateNewsResponse{
		Status: domain.Status{
			Code:    domain.Success,
			Message: domain.Success,
		},
		Id: resId,
	}, nil
}

func (ucase *NewsUseCase) AddNewsDetails(ctx context.Context, newsDetails []*domain.NewsDetails, news_id int32) (domain.CreateNewsDetailesResponse, error) {
	if newsDetails == nil {
		return domain.CreateNewsDetailesResponse{
			Status: domain.Status{
				Code:    domain.NotFound,
				Message: "empty newsDeatils list cannot be empty БЛЯДЬ!",
			},
		}, nil
	}

	err := ucase.repo.InsertIfNotExistsNewsDetails(ctx, newsDetails, news_id)
	// Ошибка запроса к базе
	if err != nil {
		return domain.CreateNewsDetailesResponse{
			Status: domain.Status{
				Code:    domain.ValidationError,
				Message: "validation",
			},
		}, nil

	}

	return domain.CreateNewsDetailesResponse{
		Status: domain.Status{
			Code:    domain.Success,
			Message: "Success",
		},
	}, nil
}
