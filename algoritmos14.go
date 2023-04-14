package main

import "fmt"

func main() {
	//Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que
	//deverão ser trocados de posição. Imprima o Slice resultante.

	//Variaveis
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var x int
	var y int

	//Mostrando o Slice
	fmt.Println(slice)

	//Solicitando os indices a serem trocados
	fmt.Println("Qual o primeiro indice a ser trocado?")
	fmt.Scan(&x)
	fmt.Println("Qual o segundo indice a ser trocado")
	fmt.Scan(&y)

	//Invertendo os valores
	slice[x], slice[y] = slice[y], slice[x]

	//imprime o slice
	fmt.Println(slice)
}
