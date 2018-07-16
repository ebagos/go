package main

import (
	"fmt"
)

func main() {
	const MLEN int = 200
	num := make([]int, MLEN)
	for i := 0; i < MLEN; i++ {
		num[i] = 0
	}
	num[0] = 1
	for i := 2; i <= 100; i++ {
		for j := 0; j < MLEN; j++ {
			num[j] *= i
		}
		for j := 0; j < MLEN-1; j++ {
			num[j+1] += num[j] / 10
			num[j] %= 10
		}
	}
	sum := 0
	for i := MLEN - 1; i >= 0; i-- {
		fmt.Print(num[i])
		sum += num[i]
	}
	fmt.Println("\nThe answer is", sum)
}
