package main

import "fmt"

func main() {
	var coluna int
	var linha int
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Informe um índice de linha (0 e 1 ), e um de coluna ( 0 a 2 ): ")
	fmt.Scan(&linha, &coluna)

	fmt.Println("A posição do array é :", array[linha][coluna])

}
