package main

import (
	"fmt"
)

func checkForPSDB(partido string) bool {
	switch {
	case partido == "PSDB":
		return true
	default:
		return false
	}
}

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
			fmt.Printf("Nosso candidato é do %s\n", pessol)

		} else if checkForPSDB(partidos[i]) {

			fmt.Printf("Encontrado um candidato do %s\n", partidos[i])
		}

	}
	numeroPrimo := 9

	if divisorDe9 := numeroPrimo; divisorDe9%3 == 0 {
		// numero primo confirmado
		fmt.Println("Yes!")
	}
}
