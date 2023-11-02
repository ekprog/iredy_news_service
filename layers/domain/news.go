package domain

import "time"

//
// MODELS
//

type NewsCard struct {
	Id        int64
	Title     string
	image     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// REPOSITORIES
type NewsRepository interface {
	FetchByPageNumber(int32) ([]*NewsCard, error)
	InsertIfNotExists(*NewsCard) error
}

//
// USE CASES
//
type NewsUseCase interface {
	GetNews(page int32) ([]*NewsCard, error)
	//AddNewsCard(newsCard NewsCard) error
}

// Response
type GetNewsResponse struct {
	News []*NewsCard
}
