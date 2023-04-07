package main

import "fmt"

func main() {
	var x int
	array := [10]int{2, 4, 5, 8, 7, 10, 21, 23, 27, 30}
	fmt.Println("Qual numero deseja procurar?")
	fmt.Scan(&x)
	var resultado int
	for _, i := range array {
		if x == i {
			resultado = 1
		} else {
			resultado = 2
		}
	}
	if resultado == 1 {
		fmt.Println("Seu numero foi encontrado!")
	} else {
		fmt.Println("Seu numero não foi encontrado!")
	}
}
