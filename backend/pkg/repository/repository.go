package repository

import (
	"github.com/jmoiron/sqlx"

	"backend/pkg/model"
)

type userRepo interface {
	GetUserInfo(userId, labId int) (model.UserRepo, error)
	InsertUserInfo(user model.UserRepo) error
	UpdateUserInfo(user model.UserRepo) error
}

type variance3LabRepo interface {
	UpdateLab3Variance(userId int, labId int, variance model.GeneratedLab3Variance) error
	GetLab3Variance(userId, labId int) (model.GeneratedLab3Variance, error)
	CheckLab3Variance(userId, labId int) error
	GetRandomLab3VarianceFromBank() (int, model.Variance, error)

	InsertLab3VarianceDB(variant model.Variance) error
}

type variance1BLabRepo interface {
	UpdateLab1BVariance(userId int, labId int, variance model.GeneratedVarianceLab1B) error
	GetLab1BVariance(userId, labId int) (model.UserVarianceLab1B, error)
	CheckLab1BVariance(userId, labId int) error
	GetRandomLab1BVarianceFromBank() (int, string, error)

	Save1BVarianceMainCriteria(userId, labId int, matrix [][]float64) error
	Save1BVarianceCriteriaMatrix(userId, labId int, index int, matrix [][]float64) error
	Save1BVarianceQualityMatrix(userId, labId int, index int, matrix [][]float64) error
}

type variance1ALabRepo interface {
	UpdateLab1AVariance(userId int, labId int, variance model.GeneratedLab1AVariance) error
	GetLab1AVariance(userId, labId int) (model.GeneratedLab1AVariance, error)
	CheckLab1AVariance(userId, labId int) error
	GetRandomLab1AVarianceFromBank() (int, model.Lab1AVariance, error)

	InsertLab1AVarianceDB(variant model.Lab1AVariance) error
}

type tokenRepo interface {
	UpdateToken(userId int, labId int, token string) error
	ClearToken(userId, labId int) error
	GetUserIdByToken(labId int, token string) (int, error)
}

type markRepo interface {
	UpdateCurrentStep(userId, labId, step int) error
	GetCurrentStep(userId, labId int) (int, error)
	IncrementMark(userId, labId, mark int) error
	GetCurrentMark(userId, labId int) (int, error)
}

type Repo struct {
	userRepo
	variance3LabRepo
	variance1ALabRepo
	variance1BLabRepo
	tokenRepo
	markRepo
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		userRepo:          NewUserRepo(db),
		tokenRepo:         NewTokenRepo(db),
		variance3LabRepo:  NewVarianceRepo(db),
		variance1ALabRepo: NewVarianceLab1ARepo(db),
		variance1BLabRepo: NewVarianceLab1BRepo(db),
		markRepo:          NewMarkRepo(db),
	}
}
