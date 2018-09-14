package main

import "fmt"

func makeDivisorList(num int) []int {
	rc := makeTrueDivisorList(num)
	rc = append(rc, num)
	return rc
}

func makeTrueDivisorList(num int) []int {
	if num < 1 {
		return []int{}
	} else if num == 1 || num == 2 || num == 3 {
		return []int{1}
	} else {
		divisorList := []int{}
		divisorList = append(divisorList, 1)
		for i := 2; i < num/2+1; i++ {
			if num%i == 0 {
				divisorList = append(divisorList, i)
			}
		}
		return divisorList
	}
}

func isFriend(num int) bool {
	a := makeTrueDivisorList(num)
	fa := 0
	for _, f := range a {
		fa += f
	}
	if fa == num {
		return false
	}
	b := makeTrueDivisorList(fa)
	fb := 0
	for _, f := range b {
		fb += f
	}
	if fb == num {
		return true
	}
	return false
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		if isFriend(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
