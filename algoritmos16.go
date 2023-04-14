package main

import "fmt"

func main() {
	//Crie um Array de inteiros com 10 elementos. Crie um novo Slice que contenha apenas os elementos pares do
	//Array original. Imprima o novo Slice.
	//Variaveis
	arr := [10]int{1, 3, 5, 6, 7, 10, 13, 14, 19, 22}
	slice := []int{}

	//separando os elementos pares
	for _, i := range arr {
		if i%2 == 0 {
			slice = append(slice, i)
		}
	}
	//imprimindo o resultado
	fmt.Println(slice)

}
