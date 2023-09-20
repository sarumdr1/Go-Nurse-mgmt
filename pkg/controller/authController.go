package controller

import (
	"encoding/json"
	"example/go-nurse-mgmt/pkg/database"
	"example/go-nurse-mgmt/pkg/helper"
	"example/go-nurse-mgmt/pkg/models"
	"log"
	"net/http"
)

type loginUserResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Define the structure that matches your JSON data
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDB()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create a new User in the database
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	db := database.GetDB()

	var requestData RequestData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	err := db.Where("email = ? AND password = ?", requestData.Email, requestData.Password).First(&user).Error

	if err != nil {
		log.Printf("User not found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		token, err := helper.CreateJWTToken(int(user.ID), user.Email)
		if err != nil {
			http.Error(w, "Error creating JWT token", http.StatusInternalServerError)
			return
		}

		// Send the token in the response
		w.WriteHeader(http.StatusOK)
		response := loginUserResponse{
			Token: token,
			Email: user.Email,
		}
		json.NewEncoder(w).Encode(response)

	}
}
