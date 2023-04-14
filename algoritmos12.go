package main

import "fmt"

// Crie um Array de inteiros com 5 elementos. Em seguida,
// crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
func main() {
	array := [5]int{1, 4, 6, 7, 9}
	var slice []int

	for _, i := range array {
		if i%3 == 0 {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)
}
