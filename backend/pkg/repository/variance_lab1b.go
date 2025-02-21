package repository

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"

	"backend/pkg/model"
)

type VarianceLab1bRepo struct {
	db *sqlx.DB
}

func NewVarianceLab1BRepo(db *sqlx.DB) *VarianceLab1bRepo {
	return &VarianceLab1bRepo{db: db}
}

func (r *VarianceLab1bRepo) UpdateLab1BVariance(userId int, labId int, variance model.GeneratedVarianceLab1B) error {
	mainCriteriaMatrix := make([][]float64, len(variance.Variance.MainCriteria))
	for i := range mainCriteriaMatrix {
		mainCriteriaMatrix[i] = make([]float64, len(variance.Variance.MainCriteria))
	}

	var matrices [][][]float64
	for i := range variance.Variance.MainCriteria {
		matrix := make([][]float64, len(variance.Variance.MainCriteria[i].Extra))
		for j := range matrix {
			matrix[j] = make([]float64, len(variance.Variance.MainCriteria[i].Extra))
		}

		matrices = append(matrices, matrix)
	}

	var qualityMatrices [][][]float64
	for i := range variance.Variance.MainCriteria {
		for j := range variance.Variance.MainCriteria[i].Extra {
			if variance.Variance.MainCriteria[i].Extra[j].IsCount {
				continue
			} else {
				matrix := make([][]float64, len(variance.Variance.Alternatives))
				for k := range matrix {
					matrix[k] = make([]float64, len(variance.Variance.Alternatives))
				}
				qualityMatrices = append(qualityMatrices, matrix)
			}
		}

	}

	userVariance := model.UserVarianceLab1B{
		Variance:              variance,
		MainCriteriaMatrix:    mainCriteriaMatrix,
		CriteriaMatrix:        matrices,
		QualityCriteriaMatrix: qualityMatrices,
	}

	jsonData, err := json.Marshal(userVariance)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET variance = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err = r.db.Exec(query, jsonData, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *VarianceLab1bRepo) GetLab1BVariance(userId, labId int) (model.UserVarianceLab1B, error) {
	var data userData
	query := fmt.Sprintf("SELECT variance FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&data, query, userId, labId); err != nil {
		return model.UserVarianceLab1B{}, err
	}

	var res model.UserVarianceLab1B
	if err := json.Unmarshal(data.Variance, &res); err != nil {
		return model.UserVarianceLab1B{}, err
	}

	return res, nil
}

func (r *VarianceLab1bRepo) Save1BVarianceMainCriteria(userId, labId int, matrix [][]float64) error {
	var data userData
	query := fmt.Sprintf("SELECT variance FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&data, query, userId, labId); err != nil {
		return err
	}

	var res model.UserVarianceLab1B
	if err := json.Unmarshal(data.Variance, &res); err != nil {
		return err
	}

	res.MainCriteriaMatrix = matrix

	jsonData, err := json.Marshal(res)
	if err != nil {
		return err
	}

	queryUpd := fmt.Sprintf("UPDATE %s SET variance = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err = r.db.Exec(queryUpd, jsonData, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *VarianceLab1bRepo) Save1BVarianceCriteriaMatrix(userId, labId int, index int, matrix [][]float64) error {
	var data userData
	query := fmt.Sprintf("SELECT variance FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&data, query, userId, labId); err != nil {
		return err
	}

	var res model.UserVarianceLab1B
	if err := json.Unmarshal(data.Variance, &res); err != nil {
		return err
	}

	if len(res.CriteriaMatrix)-1 < index {
		return fmt.Errorf("not found")
	}

	res.CriteriaMatrix[index] = matrix

	jsonData, err := json.Marshal(res)
	if err != nil {
		return err
	}

	queryUpd := fmt.Sprintf("UPDATE %s SET variance = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err = r.db.Exec(queryUpd, jsonData, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *VarianceLab1bRepo) Save1BVarianceQualityMatrix(userId, labId int, index int, matrix [][]float64) error {
	var data userData
	query := fmt.Sprintf("SELECT variance FROM %s WHERE user_id = $1 AND internal_lab_id = $2", usersTable)
	if err := r.db.Get(&data, query, userId, labId); err != nil {
		return err
	}

	var res model.UserVarianceLab1B
	if err := json.Unmarshal(data.Variance, &res); err != nil {
		return err
	}

	if len(res.QualityCriteriaMatrix)-1 < index {
		return fmt.Errorf("not found")
	}

	res.QualityCriteriaMatrix[index] = matrix

	jsonData, err := json.Marshal(res)
	if err != nil {
		return err
	}

	queryUpd := fmt.Sprintf("UPDATE %s SET variance = $1 WHERE user_id = $2 AND internal_lab_id = $3", usersTable)
	if _, err = r.db.Exec(queryUpd, jsonData, userId, labId); err != nil {
		return err
	}

	return nil
}

func (r *VarianceLab1bRepo) CheckLab1BVariance(userId, labId int) error {
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

func (r *VarianceLab1bRepo) GetRandomLab1BVarianceFromBank() (int, string, error) {
	var variantId int
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY RANDOM() LIMIT 1", bankVarianceLab1B)
	if err := r.db.Get(&variantId, query); err != nil {
		return 0, "", err
	}

	var data string
	query = fmt.Sprintf("SELECT variance FROM %s WHERE id = $1", bankVarianceLab1B)
	if err := r.db.Get(&data, query, variantId); err != nil {
		return 0, "", err
	}

	return variantId, data, nil
}
