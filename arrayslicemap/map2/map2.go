package main

import "fmt"

func main() {

	funcionario := map[string]float64{
		"Ana": 14000.40,
		"João": 4000.01,
		"Carlos": 6032.23}

	funcionario["Maria"] = 2000.23
	delete(funcionario, "José")

	for nome, salario := range funcionario{
		fmt.Printf("Salário de %s é %.2f\n", nome, salario)
	}
}
