package service

import "goapi/internal/repository"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers() []string {

	return s.repo.GetUsers()

}
