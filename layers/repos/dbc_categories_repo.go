package repos

import (
	"database/sql"
	"microservice/app/core"
	"microservice/layers/domain"
)

type DBCCategoriesRepo struct {
	log core.Logger
	db  *sql.DB
}

func NewDBCCategoriesRepo(log core.Logger, db *sql.DB) *DBCCategoriesRepo {
	return &DBCCategoriesRepo{log: log, db: db}
}

func (r *DBCCategoriesRepo) FetchNotEmptyByUserId(userId int64) ([]*domain.DBCCategory, error) {
	query := `select 
				c.id,
				c.name,
				c.created_at,
				c.updated_at,
				c.deleted_at
					from dbc_challenge_categories c
							 left join dbc_challenges dc on c.id = dc.category_id
						where c.user_id=$1 and
							  c.deleted_at is null and
							  dc.deleted_at is null
					group by (c.id, c.name, c.created_at, c.updated_at, c.deleted_at)
					having count(dc.id) > 0`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var result []*domain.DBCCategory
	for rows.Next() {
		item := &domain.DBCCategory{
			UserId: userId,
		}
		err := rows.Scan(&item.Id,
			&item.Name,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

func (r *DBCCategoriesRepo) FetchById(id int64) (*domain.DBCCategory, error) {
	var item = &domain.DBCCategory{
		Id: id,
	}
	query := `select 
    			user_id,
    			name,
    			created_at,
    			updated_at,
    			deleted_at
			from dbc_challenge_categories
			where id=$1
			limit 1`

	err := r.db.QueryRow(query, id).Scan(
		&item.UserId,
		&item.Name,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.DeletedAt)
	switch err {
	case nil:
		return item, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (r *DBCCategoriesRepo) FetchByName(userId int64, name string) (*domain.DBCCategory, error) {
	var item = &domain.DBCCategory{
		UserId: userId,
		Name:   name,
	}
	query := `select 
    			id,
    			created_at,
    			updated_at,
    			deleted_at
			from dbc_challenge_categories
			where user_id=$1 and name=$2
			limit 1`

	err := r.db.QueryRow(query, userId, name).Scan(
		&item.Id,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.DeletedAt)
	switch err {
	case nil:
		return item, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (r *DBCCategoriesRepo) Insert(item *domain.DBCCategory) error {
	query := `INSERT INTO dbc_challenge_categories (user_id, name) VALUES ($1, $2) returning id;`
	err := r.db.QueryRow(query, item.UserId, item.Name).Scan(&item.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *DBCCategoriesRepo) Update(item *domain.DBCCategory) error {
	query := `UPDATE dbc_challenge_categories 
				SET name=$2, updated_at=now()
				WHERE id=$1`
	_, err := r.db.Exec(query, item.Id, item.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *DBCCategoriesRepo) Remove(userId, id int64) error {
	query := `DELETE dbc_challenge_categories 
				SET deleted_at=now()
				WHERE id=$1 and user_id=$2`
	var str string
	err := r.db.QueryRow(query, id, userId).Scan(&str)
	if err != nil {
		return err
	}
	return nil
}
