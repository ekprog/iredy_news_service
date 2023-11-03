package repos

import (
	"database/sql"
	"fmt"
	"log"
	"microservice/app/core"
	"microservice/layers/domain"
)

type NewsRepo struct {
	log core.Logger
	db  *sql.DB
}

func NewNewsrepo(log core.Logger, db *sql.DB) *NewsRepo {
	return &NewsRepo{
		log: log,
		db:  db,
	}
}

func (r *NewsRepo) FetchByPageNumber(page int32) ([]*domain.NewsCard, error) {

	query := fmt.Sprintf("select id, title, image, type, created_at, updated_at, deleted_at from news limit %d;", page*10)

	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []*domain.NewsCard

	for rows.Next() {
		var r domain.NewsCard
		err := rows.Scan(&r.Id, &r.Title, &r.Image, &r.Type, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &r)
	}

	return result, nil

}

func (r *NewsRepo) InsertIfNotExists(card *domain.NewsCard) error {
	//TODO implement me
	panic("implement me")
}
