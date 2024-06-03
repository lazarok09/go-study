package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repository.NewUsersRepository(db)

	userID, err := userRepository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return

	}
	user.ID = userID

	responses.JSON(w, http.StatusCreated, user)

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
