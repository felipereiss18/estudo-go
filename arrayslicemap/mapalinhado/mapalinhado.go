package main

import "fmt"

func main() {

	funcionarios := map[string]map[string]float64{
		"A":{
			"Ana": 13200.32,
			"Adriana": 4223.23,
			"Alexandre": 14000.20,
		},
		"F":{
			"Felipe": 14020.21,
			"Federico": 1300.23,
		},
		"P":{
			"Paula": 1312.02,
		},
	}

	delete(funcionarios, "P")

	for grupo, funcs := range funcionarios{
		fmt.Println(grupo)
		for nome, salario := range funcs{
			fmt.Printf("Nome: %s Salário: %.2f\n", nome, salario)
		}
	}

}
