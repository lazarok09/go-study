package controllers

import (
	"encoding/json"
	"math/rand"

	"net/http"
)

const (
	Message       string = "Olá lázaro, parabéns por continuar com golang apesar do seu sono."
	FooterMessage string = "Por favor acesse a rota padrão."
)

type CustomResponse struct {
	Message       string `json:"message"`
	FooterMessage string `json:"footerMessage"`
	Id            int    `json:"id"`
}

func Store(w http.ResponseWriter, r *http.Request) {
	clientResponse := CustomResponse{Message: Message, FooterMessage: FooterMessage, Id: rand.Int()}

	data, error := json.Marshal(clientResponse)

	if error != nil {
		w.Write([]byte("An exception occurred"))
	}

	w.Write(data)
}
