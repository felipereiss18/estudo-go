package main

import (
	"github.com/felipereiss18/html"
	"time"
)

func maisRapido(url1, url2, url3 string) string {
	tit1 := html.Titulos(url1)
	tit2 := html.Titulos(url2)
	tit3 := html.Titulos(url3)

	// estrutura de controle especifica para concorrência
	select {
	case t1 := <- tit1:
		return t1
	case t2 := <- tit2:
		return t2
	case t3 := <-tit3:
		return t3
	case <- time.After(time.Millisecond * 1000):
		return "Todos perderam"
	}
}

func main() {
	ganhador := maisRapido(
			"http://www.globo.com.br",
			"http://www.amazon.com",
			"http://www.netflix.com",
		)
	println(ganhador)
}
