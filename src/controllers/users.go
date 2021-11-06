package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user-api/src/database"
	"user-api/src/models"
	"user-api/src/repositories"
	"user-api/src/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	//Validating user
	if err = user.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

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
