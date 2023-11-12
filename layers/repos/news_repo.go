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

	query := fmt.Sprintf(`SELECT id, title, image, type, created_at, updated_at, deleted_at FROM news 
						  WHERE deleted_at IS NULL and is_active is TRUE 
						  LIMIT %d OFFSET %d; `, page*10, (page-1)*10)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return []*domain.NewsCard{}, errors.Wrap(err, "Query while FetchByPageNumber")
	}
	defer rows.Close()

	var result []*domain.NewsCard

	for rows.Next() {
		var r domain.NewsCard
		err := rows.Scan(&r.Id, &r.Title, &r.Image, &r.Type, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt)
		if err != nil {
			return []*domain.NewsCard{}, errors.Wrap(err, "Scan while FetchByPageNumber")
		}
		result = append(result, &r)
	}

	return result, nil

}

func (r *NewsRepo) FetchNewsDetailsByPageNumber(ctx context.Context, page int32, news_id int32) ([]*domain.NewsDetails, error) {
	query := fmt.Sprintf(`SELECT nd.id, nd.title, nd.image, nd.type, nd.created_at, nd.updated_at FROM news_details nd
						LEFT JOIN news 
						on nd.news_id = news.id 
								WHERE nd.deleted_at is null and nd.is_active = true and nd.news_id = %d
								LIMIT %d OFFSET %d; `, news_id, page*10, (page-1)*10)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return []*domain.NewsDetails{}, errors.Wrap(err, "Query while FetchByPageNumber")
	}
	defer rows.Close()

	var result []*domain.NewsDetails

	for rows.Next() {
		var r domain.NewsDetails
		err := rows.Scan(&r.Id, &r.Title, &r.Image, &r.Type, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return []*domain.NewsDetails{}, errors.Wrap(err, "Scan while FetchByPageNumber")
		}
		result = append(result, &r)
	}

	return result, err
}

func (r *NewsRepo) InsertIfNotExistsNewsCard(ctx context.Context, card *domain.NewsCard) (int32, error) {
	type_default := "150x150" // type default описывается в миграциях БД

	// Если Title пустой, то не делаем insert
	if card.Title == "" || card.Image == "" {
		return 0, nil
	}

	// Создаём карточку новости
	query := fmt.Sprintf(`INSERT INTO news (title, image, type, is_active) 
						  VALUES ('%s', '%s', '%s', false) returning id;`,
		card.Title, card.Image, type_default)

	err := r.db.QueryRowContext(ctx, query).Scan(&card.Id)
	if err != nil {

		errors.Wrap(err, "Query while InsertIfNotExists")
		return 0, err
	}

	return card.Id, nil
}
func (r NewsRepo) InsertIfNotExistsNewsDetails(ctx context.Context, newsDetails []*domain.NewsDetails, news_id int32) error {

	for i := range newsDetails {

		// Если обязательные поля некорректны, то не делаем insert
		if newsDetails[i].Title == "" || newsDetails[i].Image == "" || newsDetails[i].NewsID <= 0 || newsDetails[i].SwipeDelay <= 0 {
			return errors.New("Hey doofus, error!")
		}

	}

	query := "INSERT INTO news_details (title, image, type, swipe_delay, news_id, is_active) VALUES "

	var is_active = true

	for i := range newsDetails {

		query += "(" +
			fmt.Sprintf("'%s','%s', '150x150', %d, %d, %t", newsDetails[i].Title, newsDetails[i].Image, newsDetails[i].SwipeDelay, newsDetails[i].NewsID, is_active)

		// Если последний элемент, то ставим скобку без запятой
		if i == len(newsDetails)-1 {
			query += ")"
		} else {
			query += "),"
		}
	}
	query += " returning id;"

	_, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "Query while InsertNews Details")
	}

	return nil
}

func (r NewsRepo) DeleteNewsCard(ctx context.Context, id int32) error {

	query := fmt.Sprintf(`update news 
						  set deleted_at = now() 
						  where id = %d`, id)

	_, err := r.db.QueryContext(ctx, query)
	if err != nil {

		errors.Wrap(err, "Query while DeleteNewsDetails")
		return err
	}

	return nil

}

func (r NewsRepo) DeleteNewsDetails(ctx context.Context, id int32) error {

	query := fmt.Sprintf(`update news_details 
						  set deleted_at = now() 
						  where id = %d`, id)

	_, err := r.db.QueryContext(ctx, query)
	if err != nil {

		errors.Wrap(err, "Query while DeleteNewsDetails")
		return err
	}

	return nil
}
