package main

import "fmt"

var daysOfMonth = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isUru(year int) int {
	rc := 0
	if year%4 == 0 {
		rc = 1
		if year%100 == 0 && year%400 != 0 {
			rc = 0
		}
	}
	return rc
}

func isNextSunday(days int) bool {
	rc := false
	if days%7 == 6 {
		rc = true
	} else {
		rc = false
	}
	return rc
}

func main() {
	count := 0
	days := 0
	for _, v := range daysOfMonth {
		days += v
	}
	days += isUru(1900)
	if isNextSunday(days) {
		count++
	}
	for year := 1901; year < 2001; year++ {
		for month := 1; month < 13; month++ {
			days += daysOfMonth[month]
			if month == 2 {
				days += isUru(year)
			}
			if isNextSunday(days) {
				if !(year == 2000 && month == 12) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
