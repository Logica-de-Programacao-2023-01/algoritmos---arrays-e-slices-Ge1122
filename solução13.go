package main

import "fmt"


func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	var num int
  
	fmt.Println("Digite um número para adicionar ao primeiro e ao último elemento da lista: ")
	fmt.Scan(&num)

	array[0] += num
	array[6] += num

	fmt.Println(array)
}
