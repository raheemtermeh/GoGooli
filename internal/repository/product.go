package repository

import (
	"context"
	"github.com/raheemtermeh/goyan-framework/internal/domain"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
}