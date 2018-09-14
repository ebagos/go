package main

import (
	"fmt"
	"sort"
)

type strSort string

func makeReverseArray(n int) []int {
	a := []int{}
	for n != 0 {
		a = append(a, n%10)
		n /= 10
	}
	return a
}

func check(n int, m int) bool {
	s1 := makeReverseArray(n)
	s2 := makeReverseArray(m)
	sort.Ints(s1)
	sort.Ints(s2)
	if len(s1) != len(s2) {
		return false
	}
	for i, a := range s1 {
		if a != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	n := 1
	for !(check(n, n*2) && check(n, n*3) && check(n, n*4) && check(n, n*5) && check(n, n*6)) {
		n++
	}
	fmt.Println(n)
}
