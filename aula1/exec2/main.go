package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	rt := chi.NewRouter()

	type user struct {
		FirstName string
		LastName  string
	}

	user1 := user{
		FirstName: "Andrea",
		LastName:  "Rivas",
	}

	rt.Get("/greetings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte(fmt.Sprintf("Hello %s %s", user1.FirstName, user1.LastName)))
	})

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte("Pong"))
	})

	http.ListenAndServe(":8081", rt)
}
