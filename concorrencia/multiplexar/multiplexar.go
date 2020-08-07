package main

import (
	"fmt"
	"github.com/felipereiss18/html"
)

func encaminhar(origem <- chan string, destino chan string)  {
	for{
		destino <- <-origem
	}
}

// Misturar mensagens num canal
func multiplexar(entrada1, entrada2 <- chan string) (ch chan string) {
	ch = make(chan string)
	go encaminhar(entrada1, ch)
	go encaminhar(entrada2, ch)

	return
}

func main() {
	ch := multiplexar(
		html.Titulos("http://www.google.com.br","http://www.udemy.com"),
		html.Titulos("http://www.amazon.com", "http://www.kabum.com.br"),
		)

	fmt.Printf("Primeiro: %s \nSegundo: %s \nTerceiro: %s \nQuarto: %s", <- ch, <- ch, <- ch, <- ch)
}
