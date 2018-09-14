package main

import (
	"fmt"
)

const maxColumn = 500
const maxTrial = 1000

func makeArray(n int) []int {
	result := make([]int, maxColumn)
	for i := 0; i < maxColumn; i++ {
		result[i] = n % 10
		n /= 10
		if n == 0 {
			break
		}
	}
	return result
}

func normalize(n []int) {
	for i := 0; i < len(n)-1; i++ {
		tmp := n[i] % 10
		n[i+1] += n[i] / 10
		n[i] = tmp
	}
}

func plus(a, b []int) []int {
	result := make([]int, maxColumn)
	for i := 0; i < maxColumn; i++ {
		result[i] = a[i] + b[i]
	}
	normalize(result)
	return result
}

func mul(a []int, n int) []int {
	result := make([]int, maxColumn)
	for i := 0; i < maxColumn; i++ {
		result[i] = a[i] * n
	}
	normalize(result)
	return result
}

func count(n []int) int {
	result := 0
	for i := maxColumn - 1; i >= 0; i-- {
		if n[i] != 0 {
			break
		}
		result++
	}
	return maxColumn - result
}

func main() {
	bunshi := makeArray(3)
	bunbo := makeArray(2)
	koeta := 0

	for i := 0; i < maxTrial; i++ {
		nbunbo := plus(bunshi, bunbo)
		nbunshi := plus(mul(bunbo, 2), bunshi)
		if count(bunbo) < count(bunshi) {
			koeta++
		}
		bunbo, bunshi = nbunbo, nbunshi
	}
	fmt.Println(koeta)
}
