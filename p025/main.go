package main

import "fmt"

func main() {
	var maxCol = 1000
	var f1 = make([]int, maxCol+1)
	var f2 = make([]int, maxCol+1)
	var f3 = make([]int, maxCol+1)

	f1[0] = 1
	f2[0] = 1
	num := 2

	for {
		for i := 0; i <= maxCol; i++ {
			f3[i] = f1[i] + f2[i]
		}
		num++
		for i := 0; i < maxCol; i++ {
			tmp := f3[i] / 10
			f3[i] %= 10
			f3[i+1] += tmp

			f1[i] = f2[i]
			f2[i] = f3[i]
		}
		if f3[999] != 0 {
			break
		}
	}
	fmt.Println(num)
}
