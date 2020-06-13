package main

import (
	"fmt"
	"math"
	"math/big"
)

func sieveOfEratosthenes(n int) []int {
	slice := make([]int, n+1)
	for i := 2; i <= n; i++ {
		slice[i] = 1
	}
	for i := 4; i <= n; i += 2 {
		slice[i] = 0
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
		if slice[i] == 0 {
			continue
		}
		// check the multiples of prime number
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			for j := 2; j*i <= n; j++ {
				slice[i*j] = 0
			}
		}
	}
	return slice
}

func main() {
	fmt.Println(sieveOfEratosthenes(50))
	fmt.Println("Prime number:")
	for i, v := range sieveOfEratosthenes(50) {
		if v == 1 {
			fmt.Println(i)
		}
	}
}
