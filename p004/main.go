package main

import (
	"fmt"
)

func kai(n int) int {
	rc := 0
	for n > 0 {
		rc = rc*10 + (n % 10)
		n /= 10
	}
	return rc
}

func main() {
	ans := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > i; j-- {
			x := i * j
			if x == kai(x) {
				if ans < x {
					fmt.Println(x)
					ans = x
				}
			}
		}
	}
	fmt.Println("The answer is", ans)
}
