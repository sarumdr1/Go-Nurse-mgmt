package controller

import (
	"encoding/json"
	"example/go-nurse-mgmt/pkg/database"
	"example/go-nurse-mgmt/pkg/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.GetDB()

	var nurse models.Nurse
	if err := json.NewDecoder(r.Body).Decode(&nurse); err != nil {
		log.Println(err)
		log.Fatal(json.NewDecoder(r.Body).Decode(&nurse))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&nurse).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created nurse as JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nurse)

}

func GetNurse(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var nurses []models.Nurse

	if err := db.Find(&nurses).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the fetched data as JSON
	json.NewEncoder(w).Encode(nurses)
}

func GetNurseById(w http.ResponseWriter, r *http.Request) {
	var nurse models.Nurse
	db := database.GetDB()

	if err := db.First(&nurse, mux.Vars(r)["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Respond with the fetched data as JSON
	json.NewEncoder(w).Encode(nurse)
}

func UpdateNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.GetDB()

	var nurse models.Nurse
	if err := db.First(&nurse, mux.Vars(r)["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewDecoder(r.Body).Decode(&nurse)
	if err := db.Save(&nurse).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(nurse)

}

func DeleteNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDB()

	var nurse models.Nurse

	if err := db.First(&nurse, mux.Vars(r)["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	if err := db.Delete(&nurse, mux.Vars(r)["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Nurse Deleted")
}
