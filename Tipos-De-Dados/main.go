package main

import (
	"errors"
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

	// Unsyned int ! NOTICE THE DIFFERENCE WITH int8
	var inteiroSemSinal uint8 = 255
	fmt.Println(inteiroSemSinal)

	// ASCII convention
	var paraCaracteresDaTabelaASCII rune = 29999

	// Alias to int8
	var inteiroDeU8bits byte = 255

	fmt.Println(inteiroSemSinal, numeroDe16Bits, numeroDe32Bits, numeroDe64Bits, numeroDe8Bits, numeroInferido, paraCaracteresDaTabelaASCII, inteiroDeU8bits)

	// suportam numeros nao inteiros
	// float32, float64
	// adicionando por inferencia
	numero := 9210.32109
	// caracteres CHAR (de um só caractere) não existem no Go.
	// Caracteres Char serão convertidos para numeros

	// Para criar strings o Go usa aspas duplas. Aspas simples é para como se fosse um Char
	// e o Char funciona como mencionamos, é um código de um caractere que referencia sua
	// posição na tabela ASCII.
	// o padrao de boleano é false
	// o padrao de string é string vazia
	// o padrao de um numero é 0
	// nill é um tipo de dado que serve de valor 0 para muitas coisas como interfaces
	// ponteiros e etc.
	// error é um tipo, errors é uma interface nativa para construção de erros em golang

	var numero1 int16 = 20202
	var numero2 float32 = 9811329312912391233912123912329210.3210393219012390123901239012390
	var numero3 float64 = 12912391233912123312123123144290292929292929229929229049219029329023902390823890239032903290123902390123901230912390123241124124231234412421422413912329210.3210393219012390123901239012390
	char := '%'
	texto := "TEXTO GRANDAO"

	var textao string
	var erro error = errors.New("errro interno")
	fmt.Println(numero, erro, numero2, numero1, numero2, numero3, char, texto, textao, erro)

}
