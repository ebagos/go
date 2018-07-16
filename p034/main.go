package main

import (
	"fmt"
)

func fill(n, c int) int {
	sum := 0
	for i := 0; i < c; i++ {
		sum = sum*10 + n
	}
	return sum
}

func factorial(n int) int {
	rc := 1
	for i := 1; i <= n; i++ {
		rc *= i
	}
	return rc
}

func limit() int {
	max := 100
	tmp := 0
	for i := 0; i < max; i++ {
		tmp = i * factorial(9)
		if fill(9, i) > tmp {
			break
		}
	}
	return tmp
}

func num2fac(n int) int {
	sum := 0
	for m := n; ; {
		x := m % 10
		sum += factorial(x)
		m /= 10
		if m == 0 {
			break
		}
	}
	return sum
}

func main() {
	sum := 0
	for i := 3; i < limit(); i++ {
		if i == num2fac(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}
