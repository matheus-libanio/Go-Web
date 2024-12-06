// internal/application/product_application.go
package application

import (
	"github.com/matheus-libanio/Go-Web/supermarket/internal/model"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/service"
)

type ProductApplication struct {
	service *service.ProductService
}

func NewProductApplication(service *service.ProductService) *ProductApplication {
	return &ProductApplication{service}
}

func (app *ProductApplication) CreateProduct(product model.Product) error {
	return app.service.CreateProduct(product)
}

func (app *ProductApplication) GetProducts() []model.Product {
	return app.service.GetProducts()
}

func (app *ProductApplication) GetProductByID(id int) (*model.Product, error) {
	return app.service.GetProductByID(id)
}

// Método para pesquisa com base em um critério, você pode adicionar mais aqui conforme necessário.
