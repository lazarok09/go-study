package main

import "fmt"

type Square struct {
	width  float64
	height float64
}

type SquareFormula interface {
	area() float64
}

func writeArea(formula SquareFormula) {
	fmt.Println(formula.area())
}
func genericFunc(unknownStuff interface{}) {
	fmt.Println(unknownStuff)
}

// como se cria um método ?

func (square Square) area() float64 {
	result := square.height * square.width

	return result
}

func main() {
	fmt.Println("Hello")

	box := Square{width: 50, height: 50}
	writeArea(box)
	genericFunc(1)
	genericFunc("Balao")
	genericFunc(float32(2912012))
}
