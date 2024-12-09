package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
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

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetById)
		r.Get("/search", handler.Search)
		r.With(checkAccessToken).Post("/", handler.Create)
		r.With(checkAccessToken).Put("/{id}", handler.Update)
		r.With(checkAccessToken).Patch("/{id}", handler.Patch)
		r.With(checkAccessToken).Delete("/{id}", handler.Delete)

	})

	log.Println("Iniciando servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

// Middleware to check Access Token
func checkAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// load
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("err loading: %v", err)
		}

		// Obtem o token do header
		token := r.Header.Get("Authorization")

		// Obtém o token esperado da variável de ambiente
		expectedToken := os.Getenv("ACCESS_TOKEN")

		// Verifica se o token da requisição é igual ao token esperado
		if token != expectedToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r) // Chama o próximo handler
	})
}
