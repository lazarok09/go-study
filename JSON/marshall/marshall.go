package marshall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Fixed bool   `json:"fixed"`
}

func RunMarshall() {

	celta := Car{
		Name:  "Celta Preto",
		Color: "Preto",
		Fixed: true,
	}
	// read the struct with marshall to create the json automatically
	bytesOfCelta, bytesOfCeltaErr := json.Marshal(celta)
	if bytesOfCeltaErr != nil {
		log.Fatal("An eror ocurred")
	}
	celtaInJSON := bytes.NewBuffer(bytesOfCelta)

	var people = map[string]string{
		"album":     "In the rainbows",
		"band":      "Radiohead",
		"firstSong": "1976",
	}
	bytesOfPeople, _ := json.Marshal(people)

	fmt.Println(celtaInJSON)
	fmt.Println(bytes.NewBuffer(bytesOfPeople))
}
