package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/application"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/model"
)

type RequestBodyProduct struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ResponseBodyProduct struct {
	Message string `json:"message"`
	Data    *Data  `json:"data,omitempty"`
	Error   bool   `json:"error"`
}

type ResponseBodyAllProducts struct {
	Message string  `json:"message"`
	Data    []*Data `json:"data,omitempty"`
	Error   bool    `json:"error"`
}

type Data struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ProductHandler struct {
	app *application.ProductApplication
}

func NewProductHandler(app *application.ProductApplication) *ProductHandler {
	return &ProductHandler{app}
}

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBodyProduct
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		code := http.StatusBadRequest
		body := &ResponseBodyProduct{
			Message: "Bad Request : invalid request body",
			Data:    nil,
			Error:   true,
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}

	product := model.Product{
		Id:           uuid.New().String(),
		Name:         reqBody.Name,
		Quantity:     reqBody.Quantity,
		Code_value:   reqBody.Code_value,
		Is_published: reqBody.Is_published,
		Expiration:   reqBody.Expiration,
		Price:        reqBody.Price,
	}

	productServ, err := p.app.Create(product)
	if err != nil {
		code := http.StatusBadRequest
		body := &ResponseBodyProduct{
			Message: "Bad Request: internal server error",
			Data:    nil,
			Error:   true,
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	dt := Data{
		Id:           productServ.Id,
		Name:         productServ.Name,
		Code_value:   productServ.Code_value,
		Is_published: productServ.Is_published,
		Expiration:   productServ.Expiration,
		Quantity:     productServ.Quantity,
		Price:        productServ.Price,
	}

	code := http.StatusCreated
	body := &ResponseBodyProduct{
		Message: "Product created",
		Data:    &dt,
		Error:   false,
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.app.GetAll()
	if err != nil {
		code := http.StatusBadRequest
		body := &ResponseBodyProduct{
			Message: "Bad Request: internal server error",
			Data:    nil,
			Error:   true,
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}

	// Construa a resposta com a struct ResponseBodyProduct
	responseData := make([]*Data, len(products)) // Alterei para um slice de Data

	for i, product := range products {
		responseData[i] = &Data{
			Id:           product.Id, // Supondo que você tem um campo Id em seu modelo
			Name:         product.Name,
			Quantity:     product.Quantity,
			Code_value:   product.Code_value,
			Is_published: product.Is_published,
			Expiration:   product.Expiration,
			Price:        product.Price,
		}
	}

	body := &ResponseBodyAllProducts{
		Message: "success to get products",
		Data:    responseData, // Atualizado para usar o Data corretamente
		Error:   false,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)

}

func (c *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	product := c.app.GetById(idStr)
	if product != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
		return
	}

	http.NotFound(w, r)
}

func (p *ProductHandler) Search(w http.ResponseWriter, r *http.Request) {
	priceStr := r.URL.Query().Get("price")
	fmt.Println("pegou:" + priceStr)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	filteredProducts, err := p.app.ServiceProducts.Search(price)
	if err != nil {
		http.NotFound(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredProducts)
}

func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	var reqBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		code := http.StatusBadRequest
		body := &ResponseBodyProduct{
			Message: "Bad Request : invalid request body",
			Data:    nil,
			Error:   true,
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}

	productServ := p.app.Update(idStr, reqBody)

	dt := Data{
		Id:           productServ.Id,
		Name:         productServ.Name,
		Code_value:   productServ.Code_value,
		Is_published: productServ.Is_published,
		Expiration:   productServ.Expiration,
		Quantity:     productServ.Quantity,
		Price:        productServ.Price,
	}

	code := http.StatusCreated
	body := &ResponseBodyProduct{
		Message: "Product created",
		Data:    &dt,
		Error:   false,
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)

}

func (p *ProductHandler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	var reqBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		code := http.StatusBadRequest
		body := &ResponseBodyProduct{
			Message: "Bad Request : invalid request body",
			Data:    nil,
			Error:   true,
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}

	product := p.app.Patch(idStr, reqBody)
	if product == nil {
		http.NotFound(w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	return

}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	product, err := p.app.Delete(idStr)
	if err != nil {
		http.NotFound(w, r)
	}

	// Return a 200 OK with the deleted product details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	return

}
