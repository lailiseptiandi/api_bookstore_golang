package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lailiseptiandi/api_golang_bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Println("Starting web on port http://localhost:9000")
	err := http.ListenAndServe("localhost:9000", r)
	log.Fatal(err)
}
