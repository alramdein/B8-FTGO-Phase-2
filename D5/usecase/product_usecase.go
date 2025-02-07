package usecase

import (
	"context"
	"fmt"

	"github.com/alramdein/template-go/model"
	"github.com/alramdein/template-go/repository"
)

// LOGIC BUSINESS

type IProductUsecase interface {
	CreateUser(cxt context.Context, product model.Product) error
}

type productUsecase struct {
	prodactRepo repository.IProductRepository
}

func NewProductUsecase(prodactRepo repository.IProductRepository) IProductUsecase {
	return &productUsecase{
		prodactRepo: prodactRepo,
	}
}

func (p *productUsecase) CreateUser(ctx context.Context, product model.Product) error {
	// validasi input
	if product.Name == "" && product.Price == 0 {
		return ErrNotFound
	}

	// misal mau insert ID pake UUID bisa disini

	err := p.prodactRepo.CreateProduct(ctx, &product)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
