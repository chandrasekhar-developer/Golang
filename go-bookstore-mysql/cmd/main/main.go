package main

import (
	"log"
	"net/http"

	"github.com/chandrasekhar-developer/go-bookstore-mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
