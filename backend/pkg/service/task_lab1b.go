package service

import (
	"context"
	"math/rand"
	"time"

	"backend/pkg/model"
	"backend/pkg/repository"
)

type tasksLab1b struct {
	repo *repository.Repo
}

func NewTaskLab1b(repo *repository.Repo) *tasksLab1b {
	return &tasksLab1b{
		repo: repo,
	}
}

func (s *tasksLab1b) UpdateUserVarianceLab1B(ctx context.Context, userId int, labId int, variance model.GeneratedVarianceLab1B) error {
	return s.repo.UpdateLab1BVariance(userId, labId, variance)
}

func (s *tasksLab1b) CheckLab1BVariance(ctx context.Context, userId, labId int) error {
	return s.repo.CheckLab1BVariance(userId, labId)
}

func (s *tasksLab1b) GetLab1BVariance(ctx context.Context, userId, labId int) (model.UserVarianceLab1B, error) {
	variance, err := s.repo.GetLab1BVariance(userId, labId)
	if err != nil {
		return model.UserVarianceLab1B{}, err
	}

	return variance, nil
}

func (s *tasksLab1b) GenerateLab1BTask(ctx context.Context, userId int) (model.GeneratedVarianceLab1B, error) {
	rand.Seed(time.Now().UnixNano())
	number, variantResponse, err := s.repo.GetRandomLab1BVarianceFromBank()
	if err != nil {
		return model.GeneratedVarianceLab1B{}, err
	}

	return model.GeneratedVarianceLab1B{
		Number: number,
		Variance: model.VarianceLab1B{
			Task: variantResponse,
		},
	}, nil
}
