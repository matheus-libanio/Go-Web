package repository

import (
	"errors"

	"github.com/matheus-libanio/Go-Web/product_crud/internal/model"
)

type RepositoryDB struct {
	DB map[string]*model.Product
}

func NewProductDB() RepositoryDB {
	return RepositoryDB{
		DB: make(map[string]*model.Product),
	}
}

func (r *RepositoryDB) Create(product model.Product) (model.Product, error) {
	id := len(r.DB) + 1
	r.DB[string(rune(id))] = &product
	product.Id = string(rune(id))
	return product, nil
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

func (r *RepositoryDB) Update(idStr string, product model.Product) *model.Product {
	UpdatedProduct := r.GetById(idStr)
	if UpdatedProduct == nil {
		return nil
	}
	UpdatedProduct.Name = product.Name
	UpdatedProduct.Quantity = product.Quantity
	UpdatedProduct.Code_value = product.Code_value
	UpdatedProduct.Is_published = product.Is_published
	UpdatedProduct.Expiration = product.Expiration
	UpdatedProduct.Price = product.Price

	return UpdatedProduct
}

func (r *RepositoryDB) Patch(idStr string, product model.Product) *model.Product {
	UpdatedProduct := r.GetById(idStr)
	if UpdatedProduct == nil {
		return nil
	}
	if product.Name != "" {
		UpdatedProduct.Name = product.Name
	}
	if product.Quantity != 0 {
		UpdatedProduct.Quantity = product.Quantity
	}
	if product.Code_value != "" {
		UpdatedProduct.Code_value = product.Code_value
	}
	if product.Is_published != false {
		UpdatedProduct.Is_published = product.Is_published
	}
	if product.Expiration != "" {
		UpdatedProduct.Expiration = product.Expiration
	}
	if product.Price != 0 {
		UpdatedProduct.Price = product.Price
	}

	return UpdatedProduct
}

func (r *RepositoryDB) GetById(idStr string) *model.Product {
	product, ok := r.DB[idStr]
	if !ok {
		return nil
	}

	return product
}
