package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lazarok09/treinandosql/database"
)

type BookResponse struct {
	Message string `json:"message"`
}

func BookControler(w http.ResponseWriter, r *http.Request) {
	connection, err := database.Connect()
	if err != nil {
		w.Write([]byte("An error ocurred when connecting the database"))

	}
	defer connection.Close()

	message := BookResponse{Message: "Olá mundo"}

	response, err := json.Marshal(message)
	if err != nil {
		w.Write([]byte("An error ocurred when parsing the message"))
	}

	w.Write(response)
}
