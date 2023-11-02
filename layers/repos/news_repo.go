package repos

import (
	"context"
	"fmt"
	"log"
	"microservice/app/core"
	"microservice/layers/domain"

	"github.com/jackc/pgx/v5"
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

func (r *NewsRepo) FetchByPageNumber(page int32) ([]*domain.NewsCard, error) {

	query := fmt.Sprintf("select * from news limit %d", page*10)

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
