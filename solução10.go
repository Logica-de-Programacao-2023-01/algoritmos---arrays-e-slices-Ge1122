package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min := lista[0]
	for _, valor := range lista {
		if valor < min {
			min = valor
		}
	}
	max := lista[0]
	for _, valor := range lista {
		if valor > max {
			max = valor
		}
	}
	fmt.Print("Valor mínimo: ", min, "\n")
	fmt.Print("Valor máximo: ", max)
}
