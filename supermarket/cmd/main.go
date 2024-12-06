package main

import (
	"log"
	"net/http"

	"github.com/matheus-libanio/Go-Web/supermarket/internal/application"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/handler"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/repository"
	"github.com/matheus-libanio/Go-Web/supermarket/internal/service"
)

func main() {

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productApplication := application.NewProductApplication(productService)
	productHandler := handler.NewProductHandler(productApplication)

	http.HandleFunc("/ping", productHandler.Ping)
	http.HandleFunc("/products", productHandler.GetProducts)
	http.HandleFunc("/products/create", productHandler.CreateProduct)
	http.HandleFunc("/products/get", productHandler.GetProductByID)
	log.Println("Iniciando servidor na porta 8080...")

	http.ListenAndServe(":8080", nil)
}
