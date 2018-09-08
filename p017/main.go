package main

import "fmt"

var numbers = map[int]string{
	1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
	6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
	11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
	16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty",
	90: "ninety", 1000: "onethousand",
}

func numberToWord(num int) string {
	v, ok := numbers[num]
	if ok == true {
		return v
	} else if num < 100 {
		a := num % 10
		b := (num / 10) * 10
		return numberToWord(b) + numberToWord(a)
	} else {
		a := num % 100
		b := num / 100
		if a == 0 {
			return numberToWord(b) + "handred"
		} else {
			return numberToWord(b) + "handredand" + numberToWord(a)
		}
	}
}

func toCharacterNum(word string) int {
	len := len(word)
	count := 0
	for _, i := range word {
		if string(i) == " " || string(i) == "-" {
			count++
		}
	}
	return len - count
}

func main() {
	words := ""
	for i := 1; i < 1001; i++ {
		words += numberToWord(i)
	}
	fmt.Println(toCharacterNum(words))
}
