package main

import (
	"fmt"
)

func main() {
	fmt.Println((""))

	// if init
	var partidos [5]string
	partidos[0] = "PT"
	partidos[1] = "PSOL"
	partidos[2] = "PSDB"
	partidos[3] = "PSL"
	partidos[4] = "PSOL"

	for i := 0; i < len(partidos); i++ {
		if pessol := partidos[i]; pessol == "PSL" {
			fmt.Printf("Nosso candidato é do PSOL %s", pessol)
			fmt.Println(" ")
		} else {
			fmt.Println("Não encontrado")
		}

	}

}
