package main

import "fmt"

type imprivel interface {
	toString() string
}

type pessoa struct {
	nome string
	sobrenome string
}

type produto struct {
	nome string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprivel)  {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.9}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Social", 82.2}
	fmt.Println(p2.toString())
	imprimir(p2)
}
