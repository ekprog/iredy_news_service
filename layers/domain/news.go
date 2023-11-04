package domain

import (
	"context"
	pb "microservice/pkg/pb/api"
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

type NewsCard struct {
	Id        int32
	Title     string
	Image     string
	Type      string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt *time.Time
}

// REPOSITORIES
type NewsRepository interface {
	FetchByPageNumber(ctx context.Context, page int32) ([]*NewsCard, error)
	InsertIfNotExists(ctx context.Context, newsCard *NewsCard) error
}

// USE CASES
type NewsUseCase interface {
	GetNews(ctx context.Context, page int32) (GetNewsResponse, error)
	AddNewsCard(ctx context.Context, newsCard NewsCard) (CreateNewsResponse, error)
}

// Response
type GetNewsResponse struct {
	Status pb.Status // ToDo: грубая ошибка. pb.* вообще никак не связано с доменом (Домен не зависит от сетевого протокола)
	News   []*NewsCard
}

type CreateNewsResponse struct {
	Status pb.Status
	Id     int32
}
