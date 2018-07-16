package main

import (
	"fmt"
)

func filter(n int) bool {
	if n%3 == 0 || n%5 == 0 {
		return true
	}
	return false
}

func fn(val int, sum int) int {
	return sum + val
}

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if filter(i) {
			sum = fn(i, sum)
		}
	}
	fmt.Println("The answer is", sum)
}
