package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paridhimaheshwari/book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
