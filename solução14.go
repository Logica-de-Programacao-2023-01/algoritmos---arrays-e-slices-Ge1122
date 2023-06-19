package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}
	var x int
  var y int
	fmt.Println("Digite o primeiro elemento:")
	fmt.Scan(&x)
	fmt.Println("Digite o segundo elemento:")
	fmt.Scan(&y)

	slice[x], slice[y] = slice[y], slice[x]

	fmt.Println(slice)
}
