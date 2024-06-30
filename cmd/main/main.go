package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utkarsh1706/Golang-BookManagementSystem/pkg/config"
	"github.com/utkarsh1706/Golang-BookManagementSystem/pkg/routes"
)

func main() {
	// Initialize MongoDB connection
	config.Connect()

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterBookStoreRoutes(r)

	// Serve HTTP using the router
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
