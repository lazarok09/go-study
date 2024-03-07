package main

import "fmt"

func treatPanicForNumber(number int) {
	// armazena o resultado utiliando o pattern 'r' do go para valores recuperados
	if r := recover(); r != nil {
		fmt.Println("Resultado do recover é: ", r)
		return
	}
	fmt.Println("Número plausível")
}

// This function may return a panic.
func checkNumber(number int) int {
	defer treatPanicForNumber(number)
	if number == 5 {
		panic("Não pode ser o valor 5!")
	}
	return number

}
func main() {
	fmt.Println("Started")
	checkNumber(5)
	fmt.Println("Ended")
}
