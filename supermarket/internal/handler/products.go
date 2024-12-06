// internal/handler/product_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/matheus-libanio/Go-Web/supermarket/internal/application"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/model"
)

type ProductHandler struct {
	app *application.ProductApplication
}

func NewProductHandler(app *application.ProductApplication) *ProductHandler {
	return &ProductHandler{app}
}

func (h *ProductHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.app.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, err := h.app.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.app.CreateProduct(product); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
