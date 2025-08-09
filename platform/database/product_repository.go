package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/raheemtermeh/goyan-framework/internal/domain"
	"github.com/raheemtermeh/goyan-framework/platform/database/sqlc_generated"
)

type productRepository struct {
	q *sqlc_generated.Queries
}

func NewProductRepository(queries *sqlc_generated.Queries) *productRepository {
	return &productRepository{q: queries}
}

func (r *productRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	dbProducts, err := r.q.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, p := range dbProducts {
		products = append(products, domain.Product{
			ID:   int(p.ID),
			Name: p.Name,
		})
	}
	return products, nil
}
func NewDBPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, connString)
}