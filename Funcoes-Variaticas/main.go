package main

import "fmt"

func calculateFCP(numbers ...int) int {
	// fundamental counting principle
	// N = n1 * n2 * n3 * ...nk
	var N int = numbers[0]

	for _, number := range numbers {
		N *= number
	}
	return N
}

func main() {
	fmt.Println(calculateFCP(1, 2, 3, 8))
}
