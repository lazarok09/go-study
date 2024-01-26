package main

import (
	"fmt"

	"github.com/go-study/lazarok09/auxiliar"
)

func main() {
	// Variávle implícita
	var nome string = "Lazaro"
	// Explicita
	var nomeExplicito = "Matheus"
	fmt.Println(nome)
	fmt.Println(nomeExplicito)

	// Iniciando duas variaveis
	// Note que nao utiliza o var quando são duas variáveis deinidas em série
	counter, reducer := 1, 2
	fmt.Println(counter, reducer)
	auxiliar.SortNumbers([10]int{1, 4, 8, 4, 2, 9, 10, 2, 3, 10})
	// Podemos trocar valores de variáveis sem precisar de auxiliar

}
