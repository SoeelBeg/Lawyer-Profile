package handler

import (
	"net/http"

	"github.com/soeel/lawyer-profile/pkg/config"
	"github.com/soeel/lawyer-profile/pkg/models"
	"github.com/soeel/lawyer-profile/pkg/routes"
)

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router := routes.RegisterRoutes()
	router.ServeHTTP(w, r)
}

func init() {
	// Connect to the database
	config.ConnectDatabase()

	// Migrate the lawyer model
	models.MigrateLawyer(config.GetDB())
}
