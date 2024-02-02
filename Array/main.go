package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World")

	//	The type [n]T is an array of n values of type T.

	var doces [4]string
	// outro jeito de escrever doces é deixando o  compilador contar o tamanho (isso nao afeta o TypeOf)
	var dieta = [...]string{"Galinha", "Arroz"}
	nomes := []string{"marmelada", "doce de leite", "goiabada", "chocolate"}

	for i := 0; i < len(doces); i++ {
		doces[i] = nomes[i]
	}
	fmt.Println(doces)
	fmt.Println("Doces é um array sim ?")

	fmt.Println(reflect.TypeOf(doces))
	fmt.Println("Nomes é um array ? não é um Slice")
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Dieta é um array ?")
	fmt.Println(reflect.TypeOf(dieta))

	// Criando um array com capacidade máxima.
	// A função make recebe um tipo, um tamanho, e uma capacidade.
	// Ela retorna um slice que aponta pra um array.
	numerosReais := make([]int8, 5, 10)

	tamanho := len(numerosReais)
	capacidadeMax := cap(numerosReais)
	fmt.Println("Tamanho (length) capacidade (cap):")
	fmt.Println(tamanho, capacidadeMax)
	fmt.Println("Tipo dos numeros reais :")
	// O tipo é int8 e não 8 posições, a capacidade permanece independente de quanto cada posição pode armazenar.
	// O int8 sugere que os itens dentro do array terão limite numérico devido a tipagem que a variável recebe que posteriormente
	// será alocada em memória. A capacidade não tem relação direta com isto, sendo possível ter um tamanho maior do que o tipo.
	numerosReais = append(numerosReais, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println(reflect.TypeOf(numerosReais))
	fmt.Println(numerosReais)
	fmt.Println(cap(numerosReais))

	// Slices também podem fazer "slicing" que seria cortar outro slice pra obter seus valores

	var numerosReaisAteCinco = numerosReais[6:11]

	fmt.Println(numerosReaisAteCinco)

}
