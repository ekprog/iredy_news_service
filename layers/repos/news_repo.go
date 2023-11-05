package repos

import (
	"context"
	"database/sql"
	"fmt"
	"microservice/app/core"
	"microservice/layers/domain"

	"github.com/pkg/errors"
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

func (r *NewsRepo) FetchByPageNumber(ctx context.Context, page int32) ([]*domain.NewsCard, error) {

	query := fmt.Sprintf("SELECT id, title, image, type, created_at, updated_at, deleted_at FROM news LIMIT %d OFFSET %d;", page*10, (page-1)*10)

	// ToDo: QueryContext вместо Query
	rows, err := r.db.Query(query)
	if err != nil {

		// ToDo: обернул, но не вернул
		errors.Wrap(err, "Query while FetchByPageNumber")
	}
	defer rows.Close()

	var result []*domain.NewsCard

	for rows.Next() {
		var r domain.NewsCard
		err := rows.Scan(&r.Id, &r.Title, &r.Image, &r.Type, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt)
		if err != nil {
			errors.Wrap(err, "Scan while FetchByPageNumber")
		}
		result = append(result, &r)
	}

	return result, nil

}

func (r *NewsRepo) InsertIfNotExists(ctx context.Context, card *domain.NewsCard) (int32, error) {
	type_default := "150x150" // type default описывается в миграциях БД

	// Если Title пустой, то не делаем insert
	if card.Title == "" {
		return 0, nil
	}

	query := fmt.Sprintf("INSERT INTO news (title, image, type) VALUES ('%s', '%s', '%s') returning id;", card.Title, card.Image, type_default)
	// ToDo: QueryRowContext вместо QueryRow
	err := r.db.QueryRow(query).Scan(&card.Id)
	if err != nil {

		errors.Wrap(err, "Query while InsertIfNotExists")
		return 0, err
	}

	return card.Id, nil
}
