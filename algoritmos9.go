package main

import "fmt"

func main() {
	var numDigitado float64
	array := [6]float64{}
	//pedindo o numero a ser adicionado
	fmt.Println("Qual numero a ser adicionado?")
	fmt.Scan(&numDigitado)
	//atribuindo valores a todos os elementos do array
	for i := range array {
		array[i] = numDigitado
	}
	fmt.Println(array)
}
