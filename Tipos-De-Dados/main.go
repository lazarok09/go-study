package main

import (
	"fmt"
)

func main() {
	// numero por inferencia de tipos usa a arquitetura do processador
	// 32 bits para computadores x86
	// 64 bits para computadores x64

	numeroInferido := 5

	var numeroDe8Bits int8 = 127
	var numeroDe16Bits int16 = 29999
	var numeroDe32Bits int32 = 1999999999
	var numeroDe64Bits int64 = 8999999999999999999

	fmt.Println(numeroInferido)
	fmt.Println(numeroDe8Bits)
	fmt.Println(numeroDe16Bits)
	fmt.Println(numeroDe32Bits)
	fmt.Println(numeroDe64Bits)

}
