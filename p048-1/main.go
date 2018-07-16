package main

/*
   数値を1桁ごとに配列に格納して計算する
*/

import "fmt"

// 桁ごとに演算した後の桁あふれを1桁上に引き継ぐ
func normalize(x []int, keta int) {
	for i := 0; i < keta-1; i++ {
		x[i+1] += x[i] / 10
		x[i] %= 10
	}
}

// 足し算
func plus(a []int, result []int, keta int) {
	for i := 0; i < keta; i++ {
		result[i] += a[i]
	}
	normalize(result, keta)
}

// べき乗
func power(x []int, n int, keta int) {
	x[0] = 1 // 掛算なので初期値を1にしておく
	// ｎにしないのはロジックを単純化するため
	for th := 0; th < n; th++ {
		for i := 0; i < keta; i++ {
			x[i] *= n
		}
		normalize(x, keta)
	}
}

// 表示用：関数化しておけばデバッグにも使える
func printout(val []int, keta int) {
	for i := keta - 2; i >= 0; i-- {
		fmt.Print(val[i])
	}
	fmt.Println("")
}

func main() {
	const MAXKETA = 11
	const MAXREP = 1000

	var result = make([]int, MAXKETA)
	for i := 0; i < MAXKETA; i++ {
		result[i] = 0
	}
	for i := 1; i <= MAXREP; i++ {
		var val = make([]int, MAXKETA)
		for j := 0; j < MAXKETA; j++ {
			val[j] = 0
		}
		power(val, i, MAXKETA)
		plus(val, result, MAXKETA)
	}
	printout(result, MAXKETA)
}
