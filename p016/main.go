package main

import (
	"fmt"
)

func main() {
	const MAX int = 500
	num := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		num[i] = 0
	}
	num[0] = 1
	for n := 1; n < 1001; n++ {
		for i := 0; i < MAX; i++ {
			num[i] *= 2
		}
		for i := 0; i < MAX-1; i++ {
			x := num[i] / 10
			num[i+1] += x
			num[i] %= 10
		}
		if n < 100 {
			fmt.Println(num[0])
		}
	}

	sum := 0
	for i := MAX - 1; i >= 0; i-- {
		fmt.Print(num[i])
		sum += num[i]
	}
	fmt.Println("\n", sum)
}
