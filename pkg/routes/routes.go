package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soeel/lawyer-profile/pkg/controller"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/lawyers", controller.CreateLawyer).Methods("POST")
	r.HandleFunc("/api/lawyers", controller.GetLawyers).Methods("GET")
	r.HandleFunc("/api/lawyers/{lawyerId:[0-9]+}", controller.GetLawyerById).Methods("GET")
	r.HandleFunc("/api/lawyers/{lawyerId:[0-9]+}", controller.UpdateLawyer).Methods("PUT")
	r.HandleFunc("/api/lawyers/{lawyerId:[0-9]+}", controller.DeleteLawyer).Methods("DELETE")

	// Serve frontend static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	return r
}
