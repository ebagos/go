package main

import (
	"fmt"
)

func collatz(n, iter int) int {
	switch {
	case n == 1:
		return iter
	case n%2 == 0:
		return collatz(n/2, iter+1)
	default:
		return collatz(3*n+1, iter+1)
	}
}

func sub1() {
	ans := 0
	max := 0
	for i := 2; i < 1000000; i++ {
		x := collatz(i, 1)
		if ans < x {
			ans = x
			max = i
		}
	}
	fmt.Println(max, ans)
}

func next(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}

func add(dict map[int]int, n1 int) map[int]int {
	n2 := next(n1)
	_, ok := dict[n2]
	if ok == false {
		dict = add(dict, n2)
	}
	dict[n1] = dict[n2] + 1
	return dict
}

func sub2() {
	dict := map[int]int{1: 1}
	maxI, max := 1, 1
	for i := 2; i < 1000000; i++ {
		_, ok := dict[i]
		if ok == false {
			dict = add(dict, i)
			if dict[i] > max {
				maxI, max = i, dict[i]
			}
		}
	}
	fmt.Println(maxI, max)
}

func main() {
	sub1()
	sub2()
}
