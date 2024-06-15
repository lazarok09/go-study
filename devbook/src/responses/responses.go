package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorShape struct {
	Error string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)

	}
}
func Error(w http.ResponseWriter, statusConde int, err error) {
	JSON(w, statusConde, ErrorShape{Error: err.Error()})
}
