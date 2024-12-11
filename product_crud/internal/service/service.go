package service

import "github.com/matheus-libanio/Go-Web/product_crud/internal/model"

type ServiceBehavior interface {
	Create(product model.Product) (model.Product, error)
	GetAll() ([]*model.Product, error)
	GetById(idStr string) *model.Product
	Search(price float64) ([]*model.Product, error)
	Update(idStr string, product model.Product) *model.Product
	Patch(idStr string, product model.Product) *model.Product
	Delete(idStr string) error
}
