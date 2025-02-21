package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TokenRepo struct {
	db *sqlx.DB
}

func NewTokenRepo(db *sqlx.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

func (r *TokenRepo) UpdateToken(userId int, labId int, token string) error {
	query := fmt.Sprintf("UPDATE %s SET token = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err := r.db.Exec(query, token, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *TokenRepo) GetUserIdByToken(labId int, token string) (int, error) {
	var userId int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE token = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&userId, query, token, labId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *TokenRepo) ClearToken(userId, labId int) error {
	query := fmt.Sprintf("UPDATE %s SET token = '' WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if _, err := r.db.Exec(query, userId, labId); err != nil {
		return err
	}

	return nil
}
