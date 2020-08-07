package main

import "fmt"

func main() {

	var notas[3] float64

	notas[0], notas[1], notas[2] = 6.5, 8.8, 9.0
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++{
		total +=notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média é %.2f", media)
}
