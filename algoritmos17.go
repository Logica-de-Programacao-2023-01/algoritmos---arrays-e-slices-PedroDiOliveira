package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // criar o Array de inteiros com 10 elementos

	soma := 0 // inicializar a variável soma como 0

	for i := 0; i < len(array); i++ { // percorrer o Array usando um loop for
		if i%2 == 0 { // verificar se a posição é par
			soma += array[i] // somar o elemento na posição par ao total
		}
	}

	fmt.Println("A soma dos elementos nas posições pares do array é:", soma) // imprimir o resultado da soma
}
