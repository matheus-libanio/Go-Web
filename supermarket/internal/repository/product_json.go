// internal/repository/product_repository.go
package repository

import (
	"sync"

	"github.com/matheus-libanio/Go-Web/supermarket/internal/model"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/utils"
)

type ProductRepository struct {
	mu       sync.RWMutex
	products []model.Product
	nextID   int
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []model.Product{},
		nextID:   1,
	}
}

func (r *ProductRepository) AddProduct(product model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.productExists(product.CodeValue) {
		return utils.ErrRequestJSONInvalid
	}

	product.ID = r.nextID
	r.nextID++
	r.products = append(r.products, product)
	return nil
}

func (r *ProductRepository) GetAllProducts() []model.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.products
}

func (r *ProductRepository) FindProductByID(id int) (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, utils.ErrRequestContentTypeNotJSON
}

func (r *ProductRepository) productExists(codeValue string) bool {
	for _, product := range r.products {
		if product.CodeValue == codeValue {
			return true
		}
	}
	return false
}

// Outros métodos como GetProductsByPriceGreaterThan, etc.
