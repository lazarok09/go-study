package unmarshall

import (
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Color string `json:"color"`
	Name  string `json:"-"`
	Year  int    `json:"year"`
}

func RunUnMarshall() {

	// he is a empty struct
	var palio Car

	// he is a json
	carINJSON := `{"color":"blue","name":"palio","year":2012}`

	if error := json.Unmarshal([]byte(carINJSON), &palio); error != nil {
		log.Fatal("An exception occurred when unmarshall the car", error)
	}
	// a json that will be mapped need to don't have different types

	carWithoutYearINJSON := `{"color":"blue","name":"palio"}`
	mappedCar := make(map[string]string)

	if error := json.Unmarshal([]byte(carWithoutYearINJSON), &mappedCar); error != nil {
		log.Fatal("An exception occurred when unmarshall the car", error)
	}

	fmt.Println(palio)
	fmt.Println(mappedCar)
}
