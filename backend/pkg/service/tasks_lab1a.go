package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"context"
	"math/rand"
	"time"
)

type tasksLab1a struct {
	repo *repository.Repo
}

func NewTaskLab1a(repo *repository.Repo) *tasksLab1a {
	return &tasksLab1a{
		repo: repo,
	}
}

func (s *tasksLab1a) UpdateUserVarianceLab1A(ctx context.Context, userId int, labId int, variance model.GeneratedLab1AVariance) error {
	return s.repo.UpdateLab1AVariance(userId, labId, variance)
}

func (s *tasksLab1a) CheckLab1AVariance(ctx context.Context, userId, labId int) error {
	return s.repo.CheckLab1AVariance(userId, labId)
}

func (s *tasksLab1a) GetLab1AVariance(ctx context.Context, userId, labId int) (model.GeneratedLab1AVariance, error) {
	variance, err := s.repo.GetLab1AVariance(userId, labId)
	if err != nil {
		return model.GeneratedLab1AVariance{}, err
	}

	return variance, nil
}

func (s *tasksLab1a) GenerateLab1ATask(ctx context.Context, userId int) (model.GeneratedLab1AVariance, error) {
	rand.Seed(time.Now().UnixNano())
	number, variantResponse, err := s.repo.GetRandomLab1AVarianceFromBank()
	if err != nil {
		return model.GeneratedLab1AVariance{}, err
	}

	return model.GeneratedLab1AVariance{
		Number:   number,
		Variance: variantResponse,
	}, nil
}
