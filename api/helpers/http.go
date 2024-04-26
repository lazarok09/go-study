package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseErrorShape struct {
	Message string
	Error   string
	Status  int
}

func ThrowParamMissing(w http.ResponseWriter, param string) {
	status := http.StatusUnprocessableEntity
	message := fmt.Sprintf("Missing the %s param", param)

	responseShape := ResponseErrorShape{Message: message, Error: "", Status: status}

	responseJSON, _ := json.Marshal(responseShape)

	w.WriteHeader(status)
	w.Write(responseJSON)
	panic(message)
}
func ThrowDBConnectionError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	responseShape := ResponseErrorShape{Message: "Seems like we're facing issues connecting the database.", Error: err.Error(), Status: status}

	responseJSON, _ := json.Marshal(responseShape)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(status)
	w.Write(responseJSON)
	panic(err)
}

// a function that receives name as being what you wanted to database prepare
func ThrowAStatmentIssue(name string) {
	message := fmt.Sprintf("An issue ocurred when you tried to: %s ", name)
	panic(message)
}

func ThrowEntityNotFounded(errorText string, w http.ResponseWriter, id uint64) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	message := fmt.Sprintf("Entity with id %d was not founded", id)
	response := ResponseErrorShape{Message: message, Error: errorText, Status: status}
	json.NewEncoder(w).Encode(response)

}
