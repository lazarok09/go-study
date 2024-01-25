package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello")
	var email string = "lazarok09@gmail.com"
	error := checkmail.ValidateFormat(email)
	if error != nil {
		fmt.Println("Email inválido")

		return
	}
	fmt.Println("Email válido")
}
