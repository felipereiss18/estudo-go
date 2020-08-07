package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // forma de while em GO, lembrando que não existe o While no go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2{
		fmt.Printf("%d", i)
	}
	
	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i+=1 {
		if i%2 == 0 {
			fmt.Print("Par ")
		}else {
			fmt.Print("Impar ")
		}
	}

	// laço infinito
	for {
		fmt.Print("\nPara Sempre")
		time.Sleep(time.Second + 5)
	}
}
