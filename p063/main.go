package main

import (
	"fmt"
)

/*
n乗すべき自然数の最大値は9なのだと仮定して、
9のn乗がn桁を超えない最大値を調べてから解答を求める
*/
const maxColumn = 200

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

func power(x, n int) []int {
	result := makeArray(x)
	for i := 1; i < n; i++ {
		for j := 0; j < maxColumn; j++ {
			result[j] *= x
		}
		normalize(result)
	}
	return result
}

func getLimit() int {
	beki := 1
	for count(power(9, beki)) >= beki {
		beki++
	}
	return beki
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
	result := 0
	for i := 1; i < 10; i++ {
		for beki := 1; beki <= getLimit(); beki++ {
			tgt := power(i, beki)
			if count(tgt) == beki {
				result++
			}
		}
	}
	fmt.Println(result)
}
