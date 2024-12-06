package storage

import (
	"encoding/json"
	"net/http"

	"github.com/matheus-libanio/Go-Web/products_crud/internal/model"
)

type Products interface {
	// GetAll returns all products
	GetAll(w http.ResponseWriter, r *http.Request)
	// GetbyID returns a product by ID
	GetByID(id int) (p *model.Product, err error)
	//Save saves a product
	Save(p *model.Product) (err error)
	// UpdateOrCreate updates or creates a product if it does not exist
	UpdateOrCreate(p *model.Product) (err error)
}

type ProductAttributes struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductsMap struct {
	db     map[int]ProductAttributes
	lastId int
}

func NewProductsMap(map[int]ProductAttributes) ProductsMap {
	return ProductsMap{
		db: make(map[int]ProductAttributes),
	}
}
func (pm *ProductsMap) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []*model.Product
	for _, product := range pm.db {
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func (pm *ProductsMap) UpdateOrCreate(p *model.Product) (err error) {
	// Serialize

	attr := ProductAttributes{Name: p.Name, Quantity: p.Quantity, CodeValue: p.CodeValue, IsPublished: p.IsPublished, Expiration: p.Expiration, Price: p.Price}

	// Update
	_, ok := pm.db[p.ID]
	switch ok {
	case true:
		pm.db[p.ID] = attr
	default:
		pm.lastId++
		pm.db[pm.lastId] = attr
	}

	return
}
