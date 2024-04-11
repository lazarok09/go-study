package unmarshall

import (
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Color string `json:"color"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
}

func RunUnMarshall() {

	// he is a empty struct
	var palio Car

	// he is a json
	carINJSON := `{"color":"blue","name":"palio","year":2012}`

	if error := json.Unmarshal([]byte(carINJSON), &palio); error != nil {
		log.Fatal("An exception occurred when unmarshall the car")
	}

	fmt.Println(palio)
}
