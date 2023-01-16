package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/", http.StripPrefix("/", http.FileServer(http.Dir(filepath.Join(".", "client", "dist")))))

	fmt.Println("litening on http://localhost:3030")
	http.ListenAndServe(":3030", router)
}
