package repository

import (
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
