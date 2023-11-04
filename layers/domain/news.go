package domain

import (
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
	FetchByPageNumber(int32) ([]*NewsCard, error)
	InsertIfNotExists(*NewsCard) error
}

// USE CASES
type NewsUseCase interface {
	GetNews(page int32) (GetNewsResponse, error)
	AddNewsCard(newsCard NewsCard) error
}

// Response
type GetNewsResponse struct {
	Status pb.Status
	News   []*NewsCard
}

type CreateNewsResponse struct {
	Status pb.Status
	Id     int32
}
