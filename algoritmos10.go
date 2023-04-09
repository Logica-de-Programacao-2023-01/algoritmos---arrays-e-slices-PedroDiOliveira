package main

import "fmt"

func main() {
	min := 99999999999
	max := -99999999999
	slice := []int{1, 4, 5, 8, 10, 15, 17, 19, 24}
	for _, x := range slice {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	fmt.Println("O valor maximo é:", max)
	fmt.Println("O valor minimo é", min)
}
