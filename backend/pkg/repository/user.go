package repository

import (
	"backend/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) UpdateUserInfo(user model.UserRepo) error {
	query := fmt.Sprintf("UPDATE %s SET is_done = $1, percentage = $2, step = $3 WHERE user_id = $4 AND internal_lab_id = $5", usersTable)
	if _, err := r.db.Exec(query, user.IsDone, user.Percentage, 0, user.UserId, user.InternalLabId); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetUserInfo(userId, labId int) (model.UserRepo, error) {
	var user model.UserRepo
	query := fmt.Sprintf("SELECT user_id, internal_lab_id, external_lab_id, is_done, percentage, token FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&user, query, userId, labId); err != nil {
		return model.UserRepo{}, err
	}

	return user, nil
}

func (r *UserRepo) InsertUserInfo(user model.UserRepo) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, internal_lab_id, external_lab_id, is_done, percentage) VALUES($1, $2, $3, $4, $5)", usersTable)

	if _, err := r.db.Exec(query, user.UserId, user.InternalLabId, user.ExternalLabId, user.IsDone, user.Percentage); err != nil {
		return err
	}

	return nil
}
