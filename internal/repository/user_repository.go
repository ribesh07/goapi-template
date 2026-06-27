package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {

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
