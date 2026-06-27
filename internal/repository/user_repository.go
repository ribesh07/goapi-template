package repository

import (
	"context"
	db "goapi/internal/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db.New(pool),
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]db.User, error) {
	return r.queries.GetUsers(ctx)
}
