package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bwm7 struct {
	
}

func (b bwm7) ligarTurbo() {
	fmt.Println("Ligando turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	bwm := bwm7{}
	bwm.ligarTurbo()
	bwm.fazerBaliza()
}
