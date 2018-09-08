package main

func isPrime(num int) bool {
	yaku := 0
	rc := false
	for dv := 1; dv < num+1; dv++ {
		if num%dv == 0 {
			yaku++
		}
	}
	if yaku == 2 {
		rc = true
	} else {
		rc = false
	}
	return rc
}

func generatePrimes(n int) []int {
	primes := []int{2}
	for i := 3; len(primes) < n; i += 2 {
		if isPrime(i) == true {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	result := generatePrimes(10001)
	println(result[len(result)-1])
}
