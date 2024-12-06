package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/application"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/handler"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/repository"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/service"
)

func main() {

	db := repository.NewProductDB()
	serv := service.NewServiceProducts(db)
	app := application.NewApplicationProducts(serv)
	handler := handler.NewProductHandler(app)
	//ctrl := controller.NewControllerProducts(serv)

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Route("/products", func(r chi.Router) {
		r.Post("/", handler.Create)
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetById)
		r.Get("/search", handler.Search)
	})
	/*
		rt.Route("/products", func(r chi.Router) {
			r.Post("/", ctrl.Create)
			r.Get("/", ctrl.GetAll)
			r.Get("/{id}", ctrl.GetById)
			r.Get("/search", ctrl.Search)
		})*/

	log.Println("Iniciando servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
