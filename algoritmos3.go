package main

import "fmt"

func main() {
	//cria um array de 4 elementos (floats)
	array := [4]float64{3.5, 2.1, 4.7, 2.5}
	produto := 1.0
	//multiplica cada termo do array entre si
	for _, valor := range array {
		produto *= valor
	}
	fmt.Println(produto)
}
