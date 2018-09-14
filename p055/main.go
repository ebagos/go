package main

import (
	"fmt"
)

const max = 50

func plus(a, b []int) []int {
	result := make([]int, max)

	for i := 0; i < max; i++ {
		result[i] = a[i] + b[i]
	}
	for i := 0; i < max-1; i++ {
		tmp := result[i] / 10
		result[i] %= 10
		result[i+1] += tmp
	}
	return result
}

func makeArray(n int) []int {
	result := make([]int, max)
	for i := 0; i < max-1; i++ {
		if n == 0 {
			break
		}
		result[i] = n % 10
		n /= 10
	}
	return result
}

func reverse(n []int) []int {
	result := make([]int, max)
	checker := max - 1
	for i := max - 1; i >= 0; i-- {
		if n[i] != 0 {
			checker = i
			break
		}
	}
	for i := 0; i <= checker; i++ {
		result[checker-i] = n[i]
	}
	return result
}

func rcheck(n []int) bool {
	checker := max - 1
	for i := max - 1; i >= 0; i-- {
		if n[i] != 0 {
			checker = i
			break
		}
	}
	tmp := reverse(n)
	for i := 0; i <= checker; i++ {
		if tmp[i] != n[i] {
			return false
		}
	}
	return true
}

func isLychrel(n int) bool {
	c := makeArray(n)
	for i := 0; i < 50; i++ {
		c = plus(c, reverse(c))
		if rcheck(c) {
			return false
		}
	}
	return true
}

func main() {
	result := []int{}
	for i := 1; i <= 10000; i++ {
		if isLychrel(i) {
			result = append(result, i)
		}
	}
	fmt.Println(result)
	fmt.Println(len(result))
}
