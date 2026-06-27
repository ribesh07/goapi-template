package service

import (
	"context"
	"goapi/internal/repository"

	db "goapi/internal/store"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers(ctx context.Context) ([]db.User, error) {
	return s.repo.GetUsers(ctx)
}
