package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shoud create a new user"))
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
