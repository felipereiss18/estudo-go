package main

import (
	"fmt"
	"time"
)

func notaParaConceito(nota float64) string{
	var n = int(nota)
	switch n {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func periodoDia() {
	t := time.Now()

	switch {
	case t.Hour() < 12: fmt.Println("Bom dia!!!")
	case t.Hour() < 18: fmt.Println("Boa tarde!!!")
	default: fmt.Println("Boa noite!!!")
	}
}

func notaParaConcceitoSemParametroSwitch(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >=8 && n < 9 :
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(4))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(-2.3))
	periodoDia()
	fmt.Println(notaParaConcceitoSemParametroSwitch(4))
	fmt.Println(notaParaConcceitoSemParametroSwitch(6.9))
	fmt.Println(notaParaConcceitoSemParametroSwitch(-2.3))
}
