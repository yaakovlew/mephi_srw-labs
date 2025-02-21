package repository

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"

	"backend/pkg/model"
)

type VarianceLab1aRepo struct {
	db *sqlx.DB
}

func NewVarianceLab1ARepo(db *sqlx.DB) *VarianceLab1aRepo {
	return &VarianceLab1aRepo{db: db}
}

func (r *VarianceLab1aRepo) UpdateLab1AVariance(userId int, labId int, variance model.GeneratedLab1AVariance) error {
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

type variance struct {
	Number   int `json:"number"`
	Variance struct {
		Matrices [][][]float64 `json:"matrices"`
	} `json:"variance"`
}

type userData struct {
	Variance []byte `db:"variance"`
}

func (r *VarianceLab1aRepo) GetLab1AVariance(userId, labId int) (model.GeneratedLab1AVariance, error) {
	var data userData
	query := fmt.Sprintf("SELECT variance FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&data, query, userId, labId); err != nil {
		return model.GeneratedLab1AVariance{}, err
	}

	var res variance
	if err := json.Unmarshal(data.Variance, &res); err != nil {
		return model.GeneratedLab1AVariance{}, err
	}

	return model.GeneratedLab1AVariance{
		Number: res.Number,
		Variance: model.Lab1AVariance{
			Matrices: res.Variance.Matrices,
		},
	}, nil
}

func (r *VarianceLab1aRepo) CheckLab1AVariance(userId, labId int) error {
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

func (r *VarianceLab1aRepo) GetRandomLab1AVarianceFromBank() (int, model.Lab1AVariance, error) {
	var variantId int
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY RANDOM() LIMIT 1", bankVarianceLab1A)
	if err := r.db.Get(&variantId, query); err != nil {
		return 0, model.Lab1AVariance{}, err
	}

	var data []byte
	query = fmt.Sprintf("SELECT variance FROM %s WHERE id = $1", bankVarianceLab1A)
	if err := r.db.Get(&data, query, variantId); err != nil {
		return 0, model.Lab1AVariance{}, err
	}

	var variance model.Lab1AVariance
	if err := json.Unmarshal(data, &variance); err != nil {
		return 0, model.Lab1AVariance{}, err
	}

	return variantId, variance, nil
}

func (r *VarianceLab1aRepo) InsertLab1AVarianceDB(variant model.Lab1AVariance) error {
	jsonData, err := json.Marshal(variant)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (variance) VALUES($1)", bankVarianceLab1A)
	if _, err := r.db.Exec(query, jsonData); err != nil {
		return err
	}

	return nil
}
