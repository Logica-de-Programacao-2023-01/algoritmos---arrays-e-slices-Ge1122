package main

import "fmt"

func main() {
	//Crie uma matriz bidimensional de inteiros com 3 linhas e 4 colunas. Inicialize cada
	//elemento com o valor da soma do índice da linha e o índice da coluna. Imprima a matriz
	//resultante.

	y := [3][4]int{{0, 1, 2, 3}, {1, 2, 3, 4}, {2, 3, 4, 5}}
	for _, y := range y {
		fmt.Println(y)
	}
}
