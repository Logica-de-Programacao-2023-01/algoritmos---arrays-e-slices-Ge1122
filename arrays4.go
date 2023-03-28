package main

import "fmt"

func main() {
	x := [6]float64{1, 2, 15, 12.6, 10.5, 5}
	var soma float64

	for _, x := range x {
		soma += x
	}

	media := soma / float64(len(x))
	fmt.Println(media)

}
