package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// ter acesso ao valor do ponteiro usa-se *
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	// para obter o endereço da variável usa-se &, por referência
	inc2(&n)
	fmt.Println(n)
}
