package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUsersRepository(db)

	userId, err := userRepository.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	var userIdStr = fmt.Sprintf("%d", userId)

	w.Write([]byte(userIdStr))

}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MAYBE"))

}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MAYBE"))

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MAYBE"))

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MAYBE"))

}
