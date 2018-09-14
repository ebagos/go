package main

import "fmt"

func ketaSquare(n int) int {
	result := 0
	for n != 0 {
		result += (n % 10) * (n % 10)
		n /= 10
	}
	return result
}

func main() {
	count := 0
	max := 9*9*7 + 1 // 9,999,999に1を足しておく
	memo := make([]int, max)
	for i := 0; i < max; i++ {
		memo[i] = -1
	}
	memo[1] = 0
	memo[89] = 1
	for i := 2; i < 10000000; i++ {
		n := i
		for {
			n = ketaSquare(n)
			if memo[n] != -1 {
				break
			}
		}
		if memo[n] == 1 {
			count++
			if i < max {
				memo[i] = 1
			}
		} else if i < max {
			memo[i] = 0
		}
	}
	fmt.Println(count)
}
