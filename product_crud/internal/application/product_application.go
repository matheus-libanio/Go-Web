package application

import (
	"github.com/matheus-libanio/Go-Web/product_crud/internal/model"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/service"
)

type ProductApplication struct {
	ServiceProducts *service.ServiceProducts
}

func NewApplicationProducts(service *service.ServiceProducts) *ProductApplication {
	return &ProductApplication{
		ServiceProducts: service,
	}
}

func (pa *ProductApplication) Create(product model.Product) (model.Product, error) {
	return pa.ServiceProducts.Create(product)
}

func (pa *ProductApplication) GetAll() ([]*model.Product, error) {
	return pa.ServiceProducts.GetAll()
}

func (pa *ProductApplication) GetById(idStr string) *model.Product {
	return pa.ServiceProducts.GetById(idStr)
}

func (pa *ProductApplication) Search(price float64) (filteredProducts []*model.Product, err error) {
	return pa.ServiceProducts.Search(price)
}

func (pa *ProductApplication) Update(idStr string, product model.Product) *model.Product {
	return pa.ServiceProducts.Update(idStr, product)
}

func (pa *ProductApplication) Patch(idStr string, product model.Product) *model.Product {
	return pa.ServiceProducts.Patch(idStr, product)
}
