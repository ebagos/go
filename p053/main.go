package main

import (
	"fmt"
	"math/big"
)

var limit = 301

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	result := 0
	for n := 1; n < 101; n++ {
		for r := 1; r <= n; r++ {
			x1 := factorial(n)
			x2 := factorial(r)
			x3 := factorial(n - r)
			x1 = x1.Div(x1, x2.Mul(x2, x3))
			x := big.NewInt(1000000)
			if x1.Cmp(x) > 0 {
				result++
			}
		}
	}
	fmt.Println(result)
}
