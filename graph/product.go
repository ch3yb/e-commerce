package graph

import (
	"context"
	"e-commerce/graph/models"
	"fmt"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input models.NewProduct) (*models.Product, error) {
	err := r.Service.CreateProduct(input)
	if err != nil {
		return nil, err
	}
	return &models.Product{ID: "1"}, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}
