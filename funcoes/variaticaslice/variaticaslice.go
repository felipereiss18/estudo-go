package main

import (
	"fmt"
)

func imprimirAprovados(aprovados ...string)  {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// funciona somente com slice
	aprovados := []string{"João", "Ana", "Maria", "Pedro", "Guilherme"}
	imprimirAprovados(aprovados...)
}
