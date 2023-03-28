package main

import "fmt"

func main() {
	//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
	x := [3]int{5, 4, 6}
	y := 0
	for _, x := range x {
		y += x
	}
	fmt.Print(y)

}
