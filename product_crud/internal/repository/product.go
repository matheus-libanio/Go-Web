package repository

import (
	"errors"

	"github.com/matheus-libanio/Go-Web/product_crud/internal/model"
)

type RepositoryDB struct {
	DB map[string]*model.Product
}

func (r *RepositoryDB) Create(product model.Product) (model.Product, error) {
	id := product.Id
	r.DB[id] = &product
	return product, nil
}

func NewProductDB() RepositoryDB {
	return RepositoryDB{
		DB: make(map[string]*model.Product),
	}
}

func (r *RepositoryDB) Delete(idStr string) error {
	// Check if the product exists
	if _, exists := r.DB[idStr]; !exists {
		return errors.New("product not found")
	}

	// Delete the product
	delete(r.DB, idStr)
	return nil
}
