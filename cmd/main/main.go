package handler

import (
	"log"
	"net/http"

	"github.com/soeel/lawyer-profile/pkg/config"
	"github.com/soeel/lawyer-profile/pkg/models"
	"github.com/soeel/lawyer-profile/pkg/routes"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	config.ConnectDatabase()

	// Migrate the lawyer model
	models.MigrateLawyer(config.GetDB())

	// Register routes
	router := routes.RegisterRoutes()

	// Serve HTTP requests
	router.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", HandleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
