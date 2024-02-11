package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice := make([]int8, 0, 2)
	slice = append(slice, 1)
	slice = append(slice, 2)
	// capacidade ainda é dois
	fmt.Println(cap(slice))
	// adicionei mais um então vamos dobrar a cap pra 4
	slice = append(slice, 3)
	// o retorno vai ser 8 por conta de otimização interna do go para tipos de int8
	// caso mudemos para int16 ou mais, o comportamento normal de dobrar persiste.
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(reflect.TypeOf(slice))

}
