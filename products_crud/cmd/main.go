package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/matheus-libanio/Go-Web/products_crud/internal/handler"
	"github.com/matheus-libanio/Go-Web/products_crud/internal/storage"
)

func main() {
	//dependencias
	db := make(map[int]storage.ProductAttributes)
	st := storage.NewProductsMap(db)
	hd := handler.NewProducts(st)

	// server
	rt := chi.NewRouter()

	rt.Route("/products", func(r chi.Router) {
		//update or create
		rt.Put("/{id}", hd.UpdateOrCreate)
	},
	)

}
