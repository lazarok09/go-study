package main

import (
	"fmt"
)

// tipo atribuido pela última tipagem
// Retorna dois tipos utilizand oos parênteses
func calculoSomaESub(a, b int8) (int8, int8) {
	var sub int8 = a - b
	var sum int8 = a + b
	return sum, sub
}

func sumAndSub(a, b int8) (sum int8, sub int8) {
	sum = a + b
	sub = a - b
	return
}

func main() {

	var funcCallBack = func() (int8, int8) {
		return calculoSomaESub(5, 8)
	}

	// Utiliza-se o underline para ocultar um valor
	//  que você não deseja em uma função que retorna dois valores
	var subtrair = func(a, b int8) int8 {
		var _, sub int8 = calculoSomaESub(a, b)
		return sub
	}

	fmt.Println(funcCallBack())
	fmt.Println(subtrair(5, 2))
	fmt.Println(sumAndSub(5, 2))
}
