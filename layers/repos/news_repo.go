package repos

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"microservice/app/core"
	"microservice/layers/domain"
)

type NewsRepo struct {
	log core.Logger
	db  *pgx.Conn
}

func NewNewsrepo(log core.Logger, db *pgx.Conn) *NewsRepo {
	return &NewsRepo{
		log: log,
		db:  db,
	}
}

func (r *NewsRepo) FetchByPageNumber(i int32) ([]*domain.NewsCard, error) {

	query := fmt.Sprintf("select * from news")

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []*domain.NewsCard

	for rows.Next() {
		var r *domain.NewsCard
		err := rows.Scan(&r.Id, &r.Title)

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, r)
	}
	return result, nil

}

func (r *NewsRepo) InsertIfNotExists(card *domain.NewsCard) error {
	//TODO implement me
	panic("implement me")
}
