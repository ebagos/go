package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//英文だからスペースが一番多いものが正解というかなりあれな根拠から解く

func main() {
	bin, err := ioutil.ReadFile("cipher.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	str := string(bin)
	str = strings.TrimRight(str, "\n")
	strs := strings.Split(str, ",")
	codes := []int{}
	for i := 0; i < len(strs); i++ {
		tmp, _ := strconv.Atoi(strs[i])
		codes = append(codes, tmp)
	}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	maxSpace := 0
	maxSeed := ""
	maxText := ""
	for _, a0 := range alphabet {
		for _, a1 := range alphabet {
			for _, a2 := range alphabet {
				seed := []int{int(a0), int(a1), int(a2)}
				text := ""
				for i, x := range codes {
					text += string(x ^ seed[i%3])
				}
				space := strings.Count(text, " ")
				if maxSpace < space {
					maxSpace = space
					maxSeed = string(a0) + string(a1) + string(a2)
					maxText = text
				}
			}
		}
	}
	fmt.Println(maxText)
	fmt.Println(maxSeed)
}
