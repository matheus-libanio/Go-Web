package service

import (
	"github.com/matheus-libanio/Go-Web/supermarket/internal/model"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) CreateProduct(product model.Product) error {
	return s.repo.AddProduct(product)
}

func (s *ProductService) GetProducts() []model.Product {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.FindProductByID(id)
}

// Outros métodos para a lógica de negócios como search, etc.
