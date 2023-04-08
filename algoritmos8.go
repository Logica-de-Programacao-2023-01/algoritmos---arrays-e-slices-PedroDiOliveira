package main

import "fmt"

func main() {
	slice := make([]string, 8)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	slice[3] = "ab"
	slice[4] = "ac"
	slice[5] = "bc"
	slice[6] = "ac"
	slice[7] = "aa"
	var resultado []string
	var valor string
	fmt.Println("Digite um valor")
	fmt.Scan(&valor)
	for _, x := range slice {
		if x != valor {
			resultado = append(resultado, x)
		}
	}
	fmt.Println(resultado)
}
