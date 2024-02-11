package main

import (
	"fmt"
)

func sleepThenPrint(i int) {
	// time.Sleep(time.Second)
	fmt.Println("Incrementando ", i)
}
func main() {

	// for classico
	for i := 0; i < 3; i++ {
		sleepThenPrint(i)
	}
	// for pelo value
	var j int

	for j < 2 {
		sleepThenPrint(j)
		j++
	}
	mapped := []string{"Milades", "FooBar", "Random"}

	// for pelo range
	for indice, name := range mapped {
		fmt.Print("Indice ", indice)
		fmt.Println(" Name", name)
	}
	// for pelo range de uma string
	for indice, letra := range "ALFABETO" {
		// O CHAR sempre vai usar o código  da tabela  ASCII
		fmt.Printf("Indice %d Código %d ASCII Letra %s  \n", indice, letra, string(letra))
	}
	user := map[string]string{
		"name":     "Lazaro",
		"lastName": "Souza",
	}
	for key, name := range user {
		fmt.Println(key, name)
	}

}
