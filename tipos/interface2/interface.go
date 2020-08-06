package main

import "fmt"

type esportivo interface {
	ligarTurbo()
	desligarTurbo()
}

type ferrari struct {
	modelo string
	turboLigado bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func (f *ferrari) desligarTurbo() {
	f.turboLigado = false
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// quando criar variável a partir de uma interace e os métodos utilizar ponteiro, precisa de criar com referência &
	var carro2 esportivo = &ferrari{"F40", true, 0}
	carro2.desligarTurbo()

	fmt.Println(carro1, carro2)
}
