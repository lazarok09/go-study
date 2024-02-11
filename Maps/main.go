package main

import (
	"fmt"
)

func main() {
	type user struct {
		id       int
		email    string
		password string
	}
	mappedUsersBySection := map[string]user{
		"compras":   {1, "Tonia", "123"},
		"estoque":   {2, "Daniel", "123"},
		"ferragens": {3, "Maurício", "123"},
	}

	fmt.Println(mappedUsersBySection)
	// delete o estoque

	delete(mappedUsersBySection, "estoque")
	fmt.Println(mappedUsersBySection)

}
