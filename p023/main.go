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

func isOver(num int) bool {
	rc := false
	a := makeTrueDivisorList(num)
	fa := 0
	for _, f := range a {
		fa += f
	}
	if fa > num {
		rc = true
	} else {
		rc = false
	}
	return rc
}

func main() {
	ar := []int{}
	dict := map[int]bool{}
	for i := 1; i <= 28123; i++ {
		rc := isOver(i)
		dict[i] = rc
		if rc {
			ar = append(ar, i)
		}
	}
	for _, x := range ar {
		for _, y := range ar {
			dict[x+y] = true
		}
	}
	sum := 0
	for i := 1; i <= 28123; i++ {
		if dict[i] == false {
			sum += i
		}
	}
	fmt.Println(sum)
}
