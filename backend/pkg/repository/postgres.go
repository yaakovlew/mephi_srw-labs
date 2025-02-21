package repository

import (
	"backend/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable        = "users_done"
	bankVarianceLab3  = "bank_variance"
	bankVarianceLab1A = "bank_variance_1a"
	bankVarianceLab1B = "bank_variance_1b"
)

func NewPostgresDB(cfg model.Config) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
