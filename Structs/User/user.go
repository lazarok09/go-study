package User

import "github.com/lazarok09/go-study/Address"

type UserStruct struct {
	Name    string
	Age     uint8
	Address Address.AddressStruct
}

func MutateName(name string, user *UserStruct) {
	user.Name = name
}

func MutateAddress(address Address.AddressStruct, user *UserStruct) {
	user.Address = address
}
