package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soeel/lawyer-profile/pkg/config"
	"github.com/soeel/lawyer-profile/pkg/models"
)

func CreateLawyer(w http.ResponseWriter, r *http.Request) {
	var lawyer models.Lawyer
	json.NewDecoder(r.Body).Decode(&lawyer)

	db := config.GetDB()
	db.Create(&lawyer)

	json.NewEncoder(w).Encode(lawyer)
}

func GetLawyers(w http.ResponseWriter, r *http.Request) {
	var lawyers []models.Lawyer

	db := config.GetDB()
	db.Find(&lawyers)

	json.NewEncoder(w).Encode(lawyers)
}

func GetLawyerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["lawyerId"], 10, 64)

	var lawyer models.Lawyer
	db := config.GetDB()
	db.First(&lawyer, id)

	json.NewEncoder(w).Encode(lawyer)
}

func UpdateLawyer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["lawyerId"], 10, 64)

	var updatedLawyer models.Lawyer
	json.NewDecoder(r.Body).Decode(&updatedLawyer)

	var lawyer models.Lawyer
	db := config.GetDB()
	db.First(&lawyer, id)

	lawyer.Name = updatedLawyer.Name
	lawyer.Specialty = updatedLawyer.Specialty
	lawyer.Email = updatedLawyer.Email
	lawyer.MobileNumber = updatedLawyer.MobileNumber
	lawyer.Address = updatedLawyer.Address

	db.Save(&lawyer)
	json.NewEncoder(w).Encode(lawyer)
}

func DeleteLawyer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["lawyerId"], 10, 64)

	var lawyer models.Lawyer
	db := config.GetDB()
	db.Delete(&lawyer, id)

	json.NewEncoder(w).Encode("Lawyer deleted successfully!")
}
