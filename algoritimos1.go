package main

import "fmt"

func main() {
	// Crie um Array de inteiros com 3 Faça um algoritmo quelementos e imprima a soma dos valores armazenados no Array.

	array := [3]int{3, 6, 9}
	var soma int
	for _, x := range array {
		soma += x
	}
	fmt.Println("A soma é:", soma)
}
