package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func calc(x int, s string) int {
	bs := []byte(s)
	sum := 0
	for i := 0; i < len(bs); i++ {
		sum += int(bs[i] - byte('A') + 1)
	}
	return x * sum
}

func main() {
	data, err := ioutil.ReadFile(`names.txt`)
	if err != nil {
		fmt.Println("Error :", err)
	}
	names := strings.Replace(string(data), "\"", "", -1)
	name := strings.Split(names, ",")
	sort.Strings(name)
	sum := 0
	for i := 0; i < len(name); i++ {
		sum += calc(i+1, name[i])
	}
	fmt.Println(sum)
}
