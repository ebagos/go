package main

/*
   数値を1桁ごとに配列に格納して計算する
*/

import "fmt"

type valarray struct {
	val    int
	result int
}

// 桁ごとに演算した後の桁あふれを1桁上に引き継ぐ
func normalize(x []valarray) {
	for i := 0; i < len(x)-1; i++ {
		x[i+1].val += x[i].val / 10
		x[i].val %= 10
		x[i+1].result += x[i].result / 10
		x[i].result %= 10
	}
	x[len(x)-1].val = 0
	x[len(x)-1].result = 0
}

// 足し算
func plus(x []valarray) {
	for i := range x {
		x[i].result += x[i].val
	}
	normalize(x)
}

// べき乗
func power(x []valarray, n int) {
	x[0].val = 1 // 掛算なので初期値を1にしておく
	// ｎにしないのはロジックを単純化するため
	for th := 0; th < n; th++ {
		for i := range x {
			x[i].val *= n
		}
		normalize(x)
	}
}

// 表示用：関数化しておけばデバッグにも使える
func printout(x []valarray) {
	for i := len(x) - 2; i >= 0; i-- {
		fmt.Print(x[i].result)
	}
	fmt.Println("")
}

func main() {
	const MAXKETA = 11
	const MAXREP = 1000

	var cdata = make([]valarray, MAXKETA)
	for i := range cdata {
		cdata[i].result = 0
	}
	for i := 1; i <= MAXREP; i++ {
		for j := range cdata {
			cdata[j].val = 0
		}
		power(cdata, i)
		plus(cdata)
	}
	printout(cdata)
}
