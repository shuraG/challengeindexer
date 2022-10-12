package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "example/indexer/domain"
)

func main() {
	fmt.Println("Hello, World!")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(""))
	})

	http.ListenAndServe(":3000", r)
}
