package main

import "fmt"

func main() {
	//perguntando o tamanho da slyce
	var y int
	fmt.Println("Qual o tamanho da slice?")
	fmt.Scan(&y)
	//criando a slyce com o tamanho referente a pregunta
	numeros := make([]int, y)
	//para cada slot, perguntando qual é o numero que substituira o 0
	for i := 0; i < y; i++ {
		fmt.Println("Digite o valor do elemento", i+1)
		fmt.Scan(&numeros[i])
	}
	var soma int
	for _, valor := range numeros {
		soma += valor
	}
	//printando a soma dos valores
	fmt.Println("A soma dos elementos da slice é", soma)
}
