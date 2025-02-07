package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alramdein/template-go/model"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	// ...
}

type productRepository struct {
	db *sql.DB
}

// dependecy injection
func NewProductRepository(db *sql.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

// CreateProduct creates a new product.
//
// The method returns an error if the execution failed.
func (p *productRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	// Exec untuk query yang gak return row (CREATE, UPDATE, DELETE)
	_, err := p.db.ExecContext(ctx, "INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
