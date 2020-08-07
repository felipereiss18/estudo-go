package main

import "fmt"

func main() {
	fmt.Println(obterResultado(6.2))
}

// Go não tem ternario
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	//return nota > 6 ? "Aprovado" : "Reprovado"
}
