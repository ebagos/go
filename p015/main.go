package main

import "fmt"

var routeNums = [][]int{}

func init() {
	tmp := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 22; i++ {
		routeNums = append(routeNums, tmp)
	}
}

func computeRouteNum(i int, j int) int {
	if i == 1 || j == 1 {
		return 1
	} else {
		return routeNums[i-1][j] + routeNums[i][j-1]
	}
}

func main() {
	for i := 1; i < 22; i++ {
		for j := 1; j < 22; j++ {
			route := computeRouteNum(i, j)
			routeNums[i][j] = route
		}
	}
	result := routeNums[21][21]
	fmt.Println(result)
}
