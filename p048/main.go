package main

import (
	"fmt"
)

const mlen = 500

var ans = make([]int, mlen)
var num = make([]int, mlen)

func powerplus(n, c int) {
	for i := 0; i < mlen; i++ {
		num[i] = 0
	}
	num[0] = n
	for i := 1; i < c; i++ {
		for j := 0; j < mlen; j++ {
			num[j] *= n
		}
		for j := 0; j < mlen-1; j++ {
			num[j+1] += num[j] / 10
			num[j] %= 10
		}
	}
	for i := 0; i < mlen; i++ {
		ans[i] += num[i]
	}
	for i := 0; i < mlen-1; i++ {
		ans[i+1] += ans[i] / 10
		ans[i] %= 10
	}
}

func main() {
	for i := 0; i < mlen; i++ {
		ans[i] = 0
	}

	for i := 1; i <= 1000; i++ {
		powerplus(i, i)
	}
	fmt.Print("The answer is ")
	for i := 9; i >= 0; i-- {
		fmt.Print(ans[i])
	}
	fmt.Println("")
}
