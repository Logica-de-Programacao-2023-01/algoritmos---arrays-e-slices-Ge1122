package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x[:3], x[4:5]...)
	fmt.Print(x)

}
