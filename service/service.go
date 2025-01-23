package service

import (
	"e-commerce/graph/models"
	"e-commerce/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"log"
)

type Service struct {
	Repository *repository.Repository
	Logger     *zap.Logger
}

func (s *Service) CreateProduct(input models.NewProduct) error {
	return s.Repository.CreateProductRepository(input.ID)

}

func NewService(l *zap.Logger) *Service {

	db, err := StartDatabase()
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}
	r := repository.NewRepository(l, db)

	l.Info("service created")
	return &Service{
		Logger:     l,
		Repository: r,
	}
}

func StartDatabase() (*pgxpool.Pool, error) {
	log.Printf("connected to postgres database \n")
	return nil, nil
}

func (s *Service) Ping() string {
	return "pong"
}
