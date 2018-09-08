package main

import (
	"fmt"
	"math"
)

/*
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 || num == 3 || num == 5 {
		return true
	}
	if num%2 == 0 || num%3 == 0 || num%5 == 0 {
		return false
	}
	prime := 7
	step := 4
	numSqrt := int(math.Sqrt(float64(num)))
	for prime <= numSqrt {
		if num%prime == 0 {
			return false
		}
		prime += step
		step = 6 - step
	}
	return true
}
*/

// かなり遅い
func makePrimeList0(num int) []int {
	result := []int{2}
	i := 3
	for i < num+1 {
		for _, p := range result {
			if i%p == 0 {
				goto skip
			}
		}
		result = append(result, i)
	skip:
		i += 2
	}
	return result
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
		for nonPrime := 2 * prime; nonPrime < num+1; nonPrime += prime {
			primeList[nonPrime] = 0
		}
	}
	result := []int{}
	for _, s := range primeList {
		if s != 0 {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	max := 2000000
	fmt.Println(makePrimeList(10))
	result := makePrimeList(max)
	sum := 0
	for _, x := range result {
		sum += x
	}
	fmt.Println(sum)
}
