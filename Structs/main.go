package main

import (
	"fmt"

	"github.com/lazarok09/go-study/Address"
	"github.com/lazarok09/go-study/User"
)

func main() {
	fmt.Println("Hello")

	// create a new user
	user := User.UserStruct{Name: "lazaro", Age: 23, Address: Address.AddressStruct{StreetName: "Rua Manuel Bom Jesus", HouseNumber: 32, ZipCode: "50943312", Country: "Amapá"}}

	// O uso do & é utilizado para passar o caminho do ponteiro user para dentro da função.
	// Originalmente user aponta * para -> UserStruct
	User.MutateAddress(Address.AddressStruct{StreetName: "Rua Manuel Bom Jesus", HouseNumber: 32, ZipCode: "50943312", Country: "Porto de Galinhas"}, &user)

	fmt.Println(user)
}
