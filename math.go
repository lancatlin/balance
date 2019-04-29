package main

func sum(numbers, pow []int) int {
	n := 1
	for i, number := range numbers {
		for j := 0; j < pow[i]; j++ {
			n *= number
		}
	}
	return n
}

func maximumCommonFactor(a, b int) int {
	c, d := a, b
	for true {
		if c >= d {
			if d == 0 {
				return c
			}
			c %= d
			c, d = d, c
		} else {
			c, d = d, c
		}
	}
	return -1 // error
}

func leastCommonMultiple(aNum, bNum int) int {
	factor := maximumCommonFactor(aNum, bNum)
	return aNum * bNum / factor
}

func primeFactorization(number int) (factors map[int]int) {
	return
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
