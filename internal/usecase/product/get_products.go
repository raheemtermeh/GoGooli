package product

import (
	"context"
	"github.com/raheemtermeh/goyan-framework/internal/domain"
	"github.com/raheemtermeh/goyan-framework/internal/repository"
)

type GetProductsUsecase struct {
	productRepo repository.ProductRepository
}

func NewGetProductsUsecase(repo repository.ProductRepository) *GetProductsUsecase {
	return &GetProductsUsecase{productRepo: repo}
}

func (uc *GetProductsUsecase) Execute(ctx context.Context) ([]domain.Product, error) {
	return uc.productRepo.GetAll(ctx)
}