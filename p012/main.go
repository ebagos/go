package main

import (
	"fmt"
	"math"
)

func trinum(num int) int {
	rc := 0
	for i := 1; i <= num; i++ {
		rc += i
	}
	return rc
}

func makePrimeList(num int) []int {
	if num < 2 {
		return []int{}
	}
	primeList := make([]int, num+1)
	for i := 0; i <= num; i++ {
		primeList[i] = i
	}
	primeList[1] = 0
	numSqrt := int(math.Sqrt(float64(num)))
	for _, prime := range primeList {
		if prime == 0 {
			continue
		}
		if prime > numSqrt {
			break
		}
		for nonPrime := prime * 2; nonPrime < num+1; nonPrime += prime {
			primeList[nonPrime] = 0
		}
	}
	result := []int{}
	for _, i := range primeList {
		if i != 0 {
			result = append(result, i)
		}
	}
	return result
}

func divisorNum(n int) int {
	num := n
	if num < 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		numSqrt := int(math.Sqrt(float64(num)))
		primeList := makePrimeList(numSqrt)
		divisor := 1
		for _, prime := range primeList {
			count := 1
			for num%prime == 0 {
				num /= prime
				count++
			}
			divisor *= count
		}
		if num != 1 {
			divisor *= 2
		}
		return divisor
	}
}

func main() {
	result := 1
	for divisorNum(trinum(result)) <= 500 {
		result++
	}
	fmt.Println(result)
}
