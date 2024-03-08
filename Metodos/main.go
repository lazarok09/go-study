package main

import (
	"fmt"
	"unicode/utf8"
)

type user struct {
	name string
	age  uint8
}

func (u user) spellLetters() []string {
	var letters []string

	for i := 0; i < len(u.name); i++ {
		runeValue, _ := utf8.DecodeRuneInString(u.name[i:])
		letters = append(letters, string(runeValue))
	}
	return letters
}

func (u user) isOld() bool {
	return u.age > 40

}

func main() {
	fmt.Println("This is example of methods.")
	// they are different from functions because are attached to a struct or interface

	argentino := user{name: "Maradona", age: 84}

	fmt.Printf("Is %s old ? %t \n", argentino.name, argentino.isOld())
	fmt.Println(argentino.spellLetters())
	fmt.Println(argentino)

}
