package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func (r *Repository) CreateProductRepository(id string) error {
	return nil
}

func NewRepository(l *zap.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{db: db, logger: l}
}
