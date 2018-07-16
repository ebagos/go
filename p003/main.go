package main

import (
	"fmt"
)

func prime_decompo(n int) []int {
	table := []int{1}
	i := 2
	for i*i <= n {
		for n%i == 0 {
			n /= i
			table = append(table, i)
		}
		i++
	}
	if n > 1 {
		table = append(table, n)
	}
	return table
}

func main() {
	pl := prime_decompo(600851475143)
	fmt.Println("The answer is", pl[len(pl)-1])
}
