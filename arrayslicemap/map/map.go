package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345667781] = "João"
	aprovados[43231232122] = "Maria"
	aprovados[31313232132] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados{
		fmt.Printf("%s (CPF: %d) ", nome, cpf)
	}

	fmt.Println(aprovados[12345667781])
	delete(aprovados, 12345667781)
	fmt.Println(aprovados[12345667781])
}
