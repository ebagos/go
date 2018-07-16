package main

import (
	"fmt"
	"math"
)

func sumpow(n int) int {
	x := float64((1 + n) * n / 2)
	return int(math.Pow(x, 2.0))
}

func powsum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2.0))
	}
	return sum
}

func main() {
	fmt.Println(powsum(10))
	fmt.Println(sumpow(10))

	fmt.Println("The answer is", sumpow(100)-powsum(100))
}
