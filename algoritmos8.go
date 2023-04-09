package main

import "fmt"

func main() {
	var x string
	slice := []string{"a", "b", "c", "ab", "bc", "aa", "cc", "bb"}
	resultado := false
	fmt.Println("Qual elemento deseja remover?")
	fmt.Scan(&x)
	for _, i := range slice {
		if x == i {
			resultado = true
		}
	}
	if resultado {
		slice = (slice:x)
	}else {
		fmt.Println("Não há um elemento similar na lista")
	}
}
