package main

import (
	"log"
	"net/http"

	"github.com/soeel/lawyer-profile/pkg/config"
	"github.com/soeel/lawyer-profile/pkg/models"
	"github.com/soeel/lawyer-profile/pkg/routes"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Migrate the lawyer model
	models.MigrateLawyer(config.GetDB())

	// Register routes
	router := routes.RegisterRoutes()

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
