package controller 

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
    "github.com/go-nurse-mgmt/pkg/models"

)

func CreateNurse(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")

	var nurse models.Nurse
    if err := json.NewDecoder(r.Body).Decode(&nurse); err != nil {
        fmt.Println(err)
		fmt.Println(json.NewDecoder(r.Body).Decode(&nurse))
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Create a new record in the database
    if err := db.Create(&nurse).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the created nurse as JSON
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(nurse)

}

func GetNurse(w http.ResponseWriter, r *http.Request) {
    var nurses []Nurse
    if err := db.Find(&nurses).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the fetched data as JSON
    json.NewEncoder(w).Encode(nurses)
}


func GetNurseById(w http.ResponseWriter, r *http.Request) {
    var nurse Nurse

    if err := db.First(&nurse,mux.Vars(r)["id"]).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Respond with the fetched data as JSON
    json.NewEncoder(w).Encode(nurse)
}

func UpdateNurse(w http.ResponseWriter, r *http.Request) {
	var nurse Nurse
    w.Header().Set("Content-Type", "application/json")
	db.First(&nurse,mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&nurse)
	db.Save(&nurse)
	json.NewEncoder(w).Encode(nurse)

}

func DeleteNurse(w http.ResponseWriter, r *http.Request){
	var nurse Nurse

	w.Header().Set("Content-Type","application/json")
	db.Delete(&nurse,mux.Vars(r)["id"])

	json.NewEncoder(w).Encode("Nurse Deleted")
}
