package controllers

import (
	"encoding/json"

	"net/http"
)

type CustomResponse struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	clientResponse := CustomResponse{Message: "Olá lázaro, parabéns por continuar com golang apesar do seu sono.", Id: 1}

	data, error := json.Marshal(clientResponse)

	if error != nil {
		w.Write([]byte("An exception occurred"))
	}

	w.Write(data)
}
