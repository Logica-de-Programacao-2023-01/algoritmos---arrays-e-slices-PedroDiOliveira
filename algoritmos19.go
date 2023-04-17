package main

import "fmt"

func main() {
	//Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro array
	//que seja a soma dos dois primeiros.
	// variaveis
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{6, 7, 8, 9, 10}

	array := []int{}

	for v := range arr1 {
		array = append(array, v)
	}
	for x := range arr2 {
		array = append(array, x)
	}
	fmt.Println(array)
}
