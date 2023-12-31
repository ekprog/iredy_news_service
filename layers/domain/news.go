package domain

import (
	"context"
	"time"
)

//
// MODELS
//

// RESPONSE CODES
const (
	Success         string = "success"
	ValidationError string = "validation_error"
	FieldRequired   string = "field_required"
	NotFound        string = "not_found"
	AlreadyExists   string = "already_exists"
	ServerError     string = "server_error"
)

type Status struct {
	Code    string
	Message string
}

// Новость, которая будет отображаться на главной
type NewsCard struct {
	Id       int32
	Title    string
	Image    string
	Type     string
	IsActive bool

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt *time.Time
}

// "Сториз", которые будут показываться
type NewsDetails struct {
	Id         int32
	Title      string
	Image      string
	Type       string
	NewsID     int32
	SwipeDelay int32 //in seconds

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt *time.Time
}

// REPOSITORIES
type NewsRepository interface {
	FetchByPageNumber(ctx context.Context, page int32) ([]*NewsCard, error)
	FetchNewsDetailsByPageNumber(ctx context.Context, page int32, news_id int32) ([]*NewsDetails, error)
	InsertIfNotExistsNewsCard(ctx context.Context, newsCard *NewsCard) (int32, error)
	InsertIfNotExistsNewsDetails(ctx context.Context, newsDetails []*NewsDetails, news_id int32) error
	DeleteNewsCard(ctx context.Context, id int32) error
	DeleteNewsDetails(ctx context.Context, id int32) error
}

// USE CASES
type NewsUseCase interface {
	GetNews(ctx context.Context, page int32) (GetNewsResponse, error)
	GetNewsDetails(ctx context.Context, page int32, news_id int32) (GetNewsDetailsResponse, error)
	AddNewsCard(ctx context.Context, newsCard NewsCard) (CreateNewsResponse, error)
	AddNewsDetails(ctx context.Context, newsDetails []*NewsDetails, news_id int32) (CreateNewsDetailesResponse, error)
	DeleteNewsCard(ctx context.Context, id int32) (Status, error)
	DeleteNewsDetails(ctx context.Context, id int32) (Status, error)
}

// Response
type GetNewsResponse struct {
	Status Status
	News   []*NewsCard
}

type GetNewsDetailsResponse struct {
	Status      Status
	NewsDetails []*NewsDetails
}

type CreateNewsResponse struct {
	Status Status
	Id     int32
}

type CreateNewsDetailesResponse struct {
	Status Status
}
