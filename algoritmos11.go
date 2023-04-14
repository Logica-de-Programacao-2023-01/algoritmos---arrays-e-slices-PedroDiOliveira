package main

import "fmt"

func main() {
	var matriz [3][2]int{
		{2,4}
		{1,8}
		{6,7}
	}
	var linha int
	var coluna int
	fmt.Println("Qual a linha desejada?")
	fmt.Scan(&linha)
	fmt.Println("Qual a coluna desejada?")
	fmt.Scan(&coluna)
	resultado := matriz[linha][coluna]
	fmt.Println(resultado)
}
