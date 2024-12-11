package repository

import "github.com/matheus-libanio/Go-Web/product_crud/internal/model"

type ProductDB interface {
	Create(product model.Product) (model.Product, error)
	Update(idStr string, product model.Product) *model.Product
	Patch(idStr string, product model.Product) *model.Product
	Delete(idStr string) error
	GetById(idStr string) *model.Product
}
