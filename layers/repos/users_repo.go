package repos

import (
	"context"
	"database/sql"
	trmsql "github.com/avito-tech/go-transaction-manager/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"microservice/app/core"
	"microservice/layers/domain"
)

type UsersRepo struct {
	log    core.Logger
	db     *sql.DB
	gormDB *gorm.DB
	getter *trmsql.CtxGetter
}

func NewUsersRepo(log core.Logger, db *sql.DB, getter *trmsql.CtxGetter, gormDB *gorm.DB) *UsersRepo {
	return &UsersRepo{
		log:    log,
		db:     db,
		getter: getter,
		gormDB: gormDB,
	}
}

func (r *UsersRepo) FetchById(id int64) (*domain.User, error) {

	query := `select 
    				score,
    				created_at, 
    				updated_at, 
    				deleted_at
				from users where id=$1 limit 1`

	user := &domain.User{Id: id}

	err := r.db.QueryRow(query, id).Scan(
		&user.Score,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt)
	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, errors.Wrap(err, "FetchById")
	}
}

func (r *UsersRepo) Exist(id int64) (bool, error) {
	query := `select id from users where id=$1 limit 1`
	err := r.db.QueryRow(query, id).Scan(&id)
	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (r *UsersRepo) InsertIfNotExists(user *domain.User) error {
	query := `INSERT INTO users (id) VALUES ($1) ON CONFLICT DO NOTHING;`
	_, err := r.db.Exec(query, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepo) Remove(id int64) error {
	query := `UPDATE users set deleted_at=now() where id=$1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepo) Update(user *domain.User) error {

	query := `UPDATE users
				SET updated_at=now()
				WHERE id=$1`

	_, err := r.db.Exec(query, user.Id, user.Score, user.ScoreDaily)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepo) AddScore(ctx context.Context, userId, score int64) error {

	query := `UPDATE users
				SET score=score+$2, updated_at=now()
				WHERE id=$1`

	_, err := r.getter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, query, userId, score)
	if err != nil {
		return err
	}
	return nil
}
