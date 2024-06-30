package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utkarsh1706/Golang-BookManagementSystem/pkg/config"
	"github.com/utkarsh1706/Golang-BookManagementSystem/pkg/routes"
)

func main() {
	config.Connect()

	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
