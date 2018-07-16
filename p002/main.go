package main

import (
	"fmt"
)

func fibo(n int) int {
	switch {
	case n < 1:
		return 0

	case n == 1:
		return 1
	case n == 2:
		return 2
	default:
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		n := fibo(i)
		if n > 4000000 {
			break
		}
		if n%2 == 0 {
			sum += n
		}
	}
	fmt.Println("The answer =", sum)
}
