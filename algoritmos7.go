package main

import "fmt"

func main() {
	slice := []int{1, 5, 9, 14, 23}
	//Perguntando o  numero a ser adicionado na slice
	var x int
	fmt.Println("Qual numero deseja adicionar?")
	fmt.Scan(&x)
	//verificando se o numero escrito ja esta presente na slice
	encontrado := false
	for _, i := range slice {
		if x == i {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Println("Seu numero ja esta presente na lista!")
	} else {
		slice = append(slice, x)
		fmt.Println(slice)
	}

}
