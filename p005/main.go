package main

import (
	"fmt"
)

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	if a != 0 && b != 0 {
		return (a * b) / gcd(a, b)
	}
	return 0
}

func main() {
	ans := 1
	for i := 1; i < 21; i++ {
		ans = lcm(ans, i)
	}
	fmt.Println(ans)
}
