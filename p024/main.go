package main

import (
	"fmt"
	"strings"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

func main() {
	const TARGET int = 1000000
	target := TARGET - 1
	ans := ""
	org := "0 1 2 3 4 5 6 7 8 9"
	sl := strings.Split(org, " ")
	for i := 9; i >= 0; i-- {
		x := factorial(i)
		n := sl[target/x]
		for j := target / x; j < len(sl)-1; j++ {
			sl[j] = sl[j+1]
		}
		target %= x
		ans += string(n)
	}
	fmt.Println("The answer is", ans)
}
