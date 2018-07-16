package main

import (
	"fmt"
	"strconv"
)

func fill(n, c int) int {
	sum := 0
	for i := 0; i < c; i++ {
		sum = sum*10 + n
	}
	return sum
}

func pow(n, x int) int {
	r := 1
	for i := 0; i < x; i++ {
		r *= n
	}
	return r
}

func limit(x int) int {
	max := 100
	tmp := 0
	for i := 0; i < max; i++ {
		tmp = i * pow(9, x)
		if fill(9, i) > tmp {
			break
		}
	}
	return tmp
}

func sub1() {
	POWER := 5
	ans := 0
	seed := make([]int, 10)
	for i := 0; i < 10; i++ {
		seed[i] = pow(i, POWER)
	}
	for i := 2; i < limit(POWER); i++ {
		s := strconv.Itoa(i)
		val := 0
		for j := 0; j < len(s); j++ {
			idx := int(s[j]) - int('0')
			val += seed[idx]
		}
		if val == i {
			fmt.Println("The condition meets with", i)
			ans += val
		}
	}
	fmt.Println("The answer is", ans)
}

func num2pow(n, c int) int {
	sum := 0
	for m := n; ; {
		x := m % 10
		sum += pow(x, c)
		m /= 10
		if m == 0 {
			break
		}
	}
	return sum
}

// Goの場合は文字列使わずに数値演算のほうが自然な感じがする
func sub2() {
	POWER := 5
	ans := 0
	for i := 2; i < limit(POWER); i++ {
		x := num2pow(i, POWER)
		if x == i {
			fmt.Println("The condition meets with", i)
			ans += x
		}
	}
	fmt.Println("The answer is", ans)
}

func main() {
	sub1()
	sub2()
}
