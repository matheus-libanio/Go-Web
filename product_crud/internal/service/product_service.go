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

func (s *ServiceProducts) Update(idStr string, product model.Product) *model.Product {
	UpdatedProduct := s.GetById(idStr)
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

func (s *ServiceProducts) Patch(idStr string, product model.Product) *model.Product {
	UpdatedProduct := s.GetById(idStr)
	if UpdatedProduct == nil {
		return nil
	}
	if product.Name != "" {
		UpdatedProduct.Name = product.Name
	}
	if product.Code_value != "" {
		UpdatedProduct.Code_value = product.Code_value
	}
	if product.Expiration != "" {
		UpdatedProduct.Expiration = product.Expiration
	}
	if product.Is_published != UpdatedProduct.Is_published {
		UpdatedProduct.Is_published = product.Is_published
	}
	if product.Quantity != 0 {
		UpdatedProduct.Quantity = product.Quantity
	}
	if product.Price != 0 {
		UpdatedProduct.Price = product.Price
	}
	return UpdatedProduct
}
