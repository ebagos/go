package main

func main() {
	//	result := []int{}
	for b := 2; b < 500; b++ {
		for a := 1; a < b; a++ {
			c := 1000 - a - b
			bb := b
			cc := c
			if c < b {
				bb, cc = c, b
			}
			if cc*cc-bb*bb-a*a == 0 {
				println(a, bb, cc)
				println(a * bb * cc)
				/*
					result = append(result, a)
					result = append(result, bb)
					result = append(result, cc)
				*/
				goto hit_exit
			}
		}
	}
hit_exit:
	//	println(result[0]*result[1]*result[2])
}
