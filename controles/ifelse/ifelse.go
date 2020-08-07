package main

import "fmt"

func imprimirResultado(nota float64) {
	// Não é necessário colocar parênteses, somente se for caso precedência entre as verificações
	if nota >= 7 {
		fmt.Println("Aprovado com a nota", nota)
	}else {
		fmt.Println("Reprovado com a nota", nota)
	}
}

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	}else if n >=8 && n < 9 {
		return "B"
	}else if n >= 5 && n < 8 {
		return "C"
	}else if n >= 3 && n < 5 {
		return "D"
	}else {
		return "E"
	}
}

func main() {
	imprimirResultado(7.5)
	imprimirResultado(6)
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(4))
	fmt.Println(notaParaConceito(2))
}
