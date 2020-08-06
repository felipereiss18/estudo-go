package main

import "fmt"

// Go consegue interpreta o contexto (escopo) lexo, ou seja, a função sabe em que local (memória) ela está
func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}
func main() {

	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
