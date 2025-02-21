package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MarkRepo struct {
	db *sqlx.DB
}

func NewMarkRepo(db *sqlx.DB) *MarkRepo {
	return &MarkRepo{db: db}
}

func (r *MarkRepo) UpdateCurrentStep(userId, labId, step int) error {
	query := fmt.Sprintf("UPDATE %s SET step = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err := r.db.Exec(query, step, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *MarkRepo) GetCurrentStep(userId, labId int) (int, error) {
	var step int
	query := fmt.Sprintf("SELECT step FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&step, query, userId, labId); err != nil {
		return 0, err
	}

	return step, nil
}

func (r *MarkRepo) IncrementMark(userId, labId, mark int) error {
	query := fmt.Sprintf("UPDATE %s SET percentage = percentage + $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err := r.db.Exec(query, mark, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *MarkRepo) GetCurrentMark(userId, labId int) (int, error) {
	var mark int
	query := fmt.Sprintf("SELECT percentage FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&mark, query, userId, labId); err != nil {
		return 0, err
	}

	return mark, nil
}
