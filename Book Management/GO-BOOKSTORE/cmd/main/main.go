package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnnyungirl/BookManagement/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
