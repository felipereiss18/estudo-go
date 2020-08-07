package main

import "fmt"

type item struct {
	produtoID int
	quantidade int
	preco float64
}

type pedido struct {
	userID int
	itens []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens{
		total += float64(item.quantidade) * item.preco
	}

	return total
}

func main() {
	p := pedido{
		userID: 1,
		itens: []item{
			{1, 23, 49},
			{2, 4, 50},
			{3, 10, 34},
		},
	}

	fmt.Printf("O valor total do pedido é R$ %.2f", p.valorTotal())
}
