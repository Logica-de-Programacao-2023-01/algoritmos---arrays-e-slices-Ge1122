package main

import "fmt"

func main() {
	//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e
	//imprima o Slice resultante.

	var x = []int{1, 2, 3, 4, 5}
	fmt.Print(x)
	x = append(x[:2], x[3:]...)
	fmt.Println(x)

}
