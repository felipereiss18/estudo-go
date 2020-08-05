package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// int para float
	fmt.Println(x / float64(y))

	// float para int
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Quando colocar o número direto ele pega o correspondente da tabela Unicode
	fmt.Println("Teste", string(97))

	// int para String
	fmt.Println("Teste", strconv.Itoa(123))

	// string para int
	// quando colocar o _ é para dizer que está ignorando o retorno respectivo da função
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	// string para boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

}
