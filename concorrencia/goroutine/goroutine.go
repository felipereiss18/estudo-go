package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)

	go fale("Maria", "Agora simm!!", 10)
	fale("João", "Tudo faz sentido!!", 5)
}
