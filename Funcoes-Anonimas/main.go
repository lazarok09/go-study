package main

import "fmt"

func main() {

	fmt.Println("Chamadas anonimas")

	// exemplo 1
	func() {
		fmt.Print("CHAMADA ANONIMA")
	}()
	func(text string) {
		fmt.Printf("CHAMADA ANONIMA com parametro %s \n", text)
	}("Parametro especial")

	var result = func(text string) string {
		return "CHAMADA ANONIMA com parametro " + text
	}("Parametro especial")
	fmt.Print(result)
}
