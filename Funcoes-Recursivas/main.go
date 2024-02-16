package main

import "fmt"

// 1 1 2 3  5 8 13 21
func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)

}

func main() {
	fmt.Println("Fibonnaci: ")
	fmt.Println(fibonacci(8))
}
