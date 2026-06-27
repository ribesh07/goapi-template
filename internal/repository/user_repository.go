package repository

import "goapi/internal/database"

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUsers() []string {

	return []string{
		"Guru",
		"Ram",
		"Hari",
	}
}
