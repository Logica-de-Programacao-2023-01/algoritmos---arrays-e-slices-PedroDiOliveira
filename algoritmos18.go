// Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.
package main

import "fmt"

// Função para verificar se um número é primo
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n int
	fmt.Println("Numero de valores primos a ser verificados?")
	fmt.Scanln(&n)

	primes := []int{}
	i := 2
	for len(primes) < n {
		if isPrime(i) {
			primes = append(primes, i)
		}
		i++
	}

	fmt.Println(primes)
}
