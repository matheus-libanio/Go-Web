package service

import (
	"errors"

	"github.com/matheus-libanio/Go-Web/product_crud/internal/model"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/repository"
)

type ServiceProducts struct {
	Storage repository.RepositoryDB
}

func (s *ServiceProducts) Create(product model.Product) (model.Product, error) {
	product, err := s.Storage.Create(product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func NewServiceProducts(storage repository.RepositoryDB) *ServiceProducts {
	return &ServiceProducts{
		Storage: storage,
	}
}

func (s *ServiceProducts) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	for _, product := range s.Storage.DB {
		products = append(products, product)
	}
	return products, nil
}

func (s *ServiceProducts) GetById(idStr string) *model.Product {
	return s.Storage.DB[idStr]
}

func (s *ServiceProducts) Search(price float64) (filteredProducts []*model.Product, err error) {
	for _, product := range s.Storage.DB {
		if product.Price > price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	if filteredProducts == nil {
		return filteredProducts, errors.New("Product not found")
	}
	return filteredProducts, nil
}
