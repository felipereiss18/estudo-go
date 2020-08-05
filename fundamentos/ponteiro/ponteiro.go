package main

import "fmt"

func main() {

	i := 1

	// Go não tem aritméticas de ponteiros
	// p++
	var p *int = nil

	p = &i // pegando o dendereço da variável
	*p++ // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, &i)

}
