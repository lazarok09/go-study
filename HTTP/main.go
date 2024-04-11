package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomResponse struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		clientResponse := CustomResponse{Message: "Olá lázaro, parabéns por continuar com golang apesar do seu sono.", Id: 1}

		data, error := json.Marshal(clientResponse)

		if error != nil {
			w.Write([]byte("An exception occured"))
		}

		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
