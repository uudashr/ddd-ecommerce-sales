package mysql

import (
	"database/sql"
	"errors"

	"github.com/uudashr/ddd-ecommerce-sales/internal/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) (*ProductRepository, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	return &ProductRepository{db}, nil
}

func (r *ProductRepository) Store(prod *product.Product) error {
	// TODO: implement this
	panic("not implemented") // implementation should be quite straigt forward
}

func (r *ProductRepository) ProductByID(id string) (*product.Product, error) {
	// TODO: implement this
	panic("not implemented") // implementation should be quite straigt forward
}

func (r *ProductRepository) Update(prod *product.Product) error {
	// TODO: implement this
	panic("not implemented") // implementation should be quite straigt forward
}
