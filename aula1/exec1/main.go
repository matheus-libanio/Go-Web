package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	rt := chi.NewRouter()

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte("Pong"))
	})

	http.ListenAndServe(":8081", rt)

}
