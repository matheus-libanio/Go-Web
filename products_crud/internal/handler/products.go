// internal/handler/product_handler.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/matheus-libanio/Go-Web/products_crud/internal/model"
	"github.com/matheus-libanio/Go-Web/products_crud/internal/storage"
	"github.com/matheus-libanio/Go-Web/products_crud/platform/web/request"
	"github.com/matheus-libanio/Go-Web/products_crud/platform/web/response"
)

// Products is a  struct that contains the methods of a handler of products
type Products struct {
	// st is the interface  Products for storage operations
	st storage.Products
}

type RequestBodyUpdateOrCreate struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func NewProducts(storage storage.Products) *Products {
	return &Products{
		st: storage,
	}
}

// UpdateOrCreate updates or creates a product if it does not exist, by id
func (p *Products) UpdateOrCreate(w http.ResponseWriter, r *http.Request) {
	// request
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusBadRequest
		body := map[string]any{"message": "invalid id", "data": nil}

		response.JSON(w, code, body)
		return
	}

	var reqBody RequestBodyUpdateOrCreate
	if err := request.JSON(r, &reqBody); err != nil {
		code := http.StatusBadRequest
		body := map[string]any{"message": "invalid request body", "data": nil}

		response.JSON(w, code, body)
		return
	}

	// desserialize
	pr := model.Product{
		ID:          id,
		Name:        reqBody.Name,
		Quantity:    reqBody.Quantity,
		CodeValue:   reqBody.CodeValue,
		IsPublished: reqBody.IsPublished,
		Expiration:  reqBody.Expiration,
		Price:       reqBody.Price,
	}

	// Update or create
	if err := p.st.UpdateOrCreate(&pr); err != nil {
		code := http.StatusInternalServerError
		body := map[string]any{"message": "internal server error", "data": nil}
		response.JSON(w, code, body)
	}

	//response
	code := http.StatusOK
	body := map[string]any{"message": "product updated or created", "data": nil}

	response.JSON(w, code, body)
}
