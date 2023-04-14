package main

import "fmt"

func main() {
	// Criando o array bidimensional
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Solicitando ao usuário os índices de linha e coluna
	var linha, coluna int
	fmt.Print("Informe o índice da linha: ")
	fmt.Scan(&linha)
	fmt.Print("Informe o índice da coluna: ")
	fmt.Scan(&coluna)

	// Imprimindo o valor armazenado na posição informada
	fmt.Printf("O valor armazenado na posição [%d][%d] é %d\n", linha, coluna, matriz[linha][coluna])
}
