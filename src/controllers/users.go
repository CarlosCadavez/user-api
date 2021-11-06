package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"user-api/src/database"
	"user-api/src/models"
	"user-api/src/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserted: %d", userId)))

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shoud retrieve a list of users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shoud retrieve a specific user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shoud update an user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shoud delete an user"))
}
