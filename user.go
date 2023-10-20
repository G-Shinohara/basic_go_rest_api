package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var users []User
	DB.Find(&users)

	json.NewEncoder(writer).Encode(users)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	var user User
	DB.First(&user, params["id"])

	json.NewEncoder(writer).Encode(user)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user User
	json.NewDecoder(request.Body).Decode(&user)

	DB.Create(&user)

	json.NewEncoder(writer).Encode(user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	var user User
	DB.First(&user, params["id"])

	json.NewDecoder(request.Body).Decode(&user)

	DB.Save(&user)

	json.NewEncoder(writer).Encode(&user)
}

func DeleteUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	var user User

	DB.Delete(&user, params["id"])

	json.NewEncoder(writer).Encode("The user was deleted")
}
