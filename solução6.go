package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	encontrado := false
	for _, valor := range array {

		if numero == valor {
			encontrado = true
			break
		}
	}
	if encontrado  {
		fmt.Print("O seu número está no Array!")
	} else {
		fmt.Print("O seu número não esta a Array!")
	}

}
