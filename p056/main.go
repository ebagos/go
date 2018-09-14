package main

import (
	"fmt"
	"math/big"
)

func main() {
	result := big.NewInt(int64(0))
	for i := int64(1); i < int64(100); i++ {
		for j := int64(1); j < int64(100); j++ {
			a := big.NewInt(i)
			b := big.NewInt(j)
			a.Exp(a, b, nil)
			sum := big.NewInt(int64(0))
			for a.Cmp(big.NewInt(int64(0))) == 1 {
				m := big.NewInt(int64(1))
				_, nn := a.DivMod(a, big.NewInt(int64(10)), m)
				sum.Add(sum, nn)
			}
			if result.Cmp(sum) == -1 {
				result = sum
			}
		}
	}
	fmt.Println(result)
}
