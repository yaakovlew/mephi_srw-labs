package repository

import (
	"backend/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VarianceLab3Repo struct {
	db *sqlx.DB
}

func NewVarianceRepo(db *sqlx.DB) *VarianceLab3Repo {
	return &VarianceLab3Repo{db: db}
}

func (r *VarianceLab3Repo) UpdateLab3Variance(userId int, labId int, variance model.GeneratedLab3Variance) error {
	jsonData, err := json.Marshal(variance)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET variance = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err = r.db.Exec(query, jsonData, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *VarianceLab3Repo) GetLab3Variance(userId, labId int) (model.GeneratedLab3Variance, error) {
	var data []byte
	query := fmt.Sprintf("SELECT variance FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&data, query, userId, labId); err != nil {
		return model.GeneratedLab3Variance{}, err
	}

	var variance model.GeneratedLab3Variance
	if err := json.Unmarshal(data, &variance); err != nil {
		return model.GeneratedLab3Variance{}, err
	}

	return variance, nil
}

func (r *VarianceLab3Repo) CheckLab3Variance(userId, labId int) error {
	var user int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&user, query, userId, labId); err != nil {
		return err
	}

	if user == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *VarianceLab3Repo) GetRandomLab3VarianceFromBank() (int, model.Variance, error) {
	var variantId int
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY RANDOM() LIMIT 1", bankVarianceLab3)
	if err := r.db.Get(&variantId, query); err != nil {
		return 0, model.Variance{}, err
	}

	var data []byte
	query = fmt.Sprintf("SELECT variance FROM %s WHERE id = $1", bankVarianceLab3)
	if err := r.db.Get(&data, query, variantId); err != nil {
		return 0, model.Variance{}, err
	}

	var variance model.Variance
	if err := json.Unmarshal(data, &variance); err != nil {
		return 0, model.Variance{}, err
	}

	return variantId, variance, nil
}

func (r *VarianceLab3Repo) InsertLab3VarianceDB(variant model.Variance) error {
	jsonData, err := json.Marshal(variant)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (variance) VALUES($1)", bankVarianceLab3)
	if _, err := r.db.Exec(query, jsonData); err != nil {
		return err
	}

	return nil
}
