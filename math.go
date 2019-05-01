package balance

import ()

func sum(numbers, pow []int) uint64 {
	var n uint64 = 1
	for i, number := range numbers {
		for j := 0; j < pow[i]; j++ {
			n *= uint64(number)
		}
	}
	return n
}

func maximumCommonFactor(a, b uint64) uint64 {
	c, d := a, b
	for {
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
}

func leastCommonMultiple(aNum, bNum uint64) uint64 {
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
