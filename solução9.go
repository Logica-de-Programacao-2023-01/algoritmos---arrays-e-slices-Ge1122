package main

import "fmt"

func main() {

	var Array = [6]float64{3.9, 6.1, 5.7, 2.6, 0.9, 2.5}
	var valor float64

	fmt.Print("Informe um número: ")
	fmt.Scan(&valor)

	for i := 0; i < len(Array); i++ {
		Array[i] += valor
	}

	fmt.Println("A array agora fica : ", Array)
}
