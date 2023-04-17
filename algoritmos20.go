package main

import "fmt"

func main() {
	// lendo o array de inteiros
	var nums []int
	fmt.Println("Digite os números do array separados por espaço:")
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		nums = append(nums, num)
	}

	// verificando se o array está ordenado em ordem crescente
	ordenado := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			ordenado = false
			break
		}
	}

	// exibindo o resultado
	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente")
	}
}
