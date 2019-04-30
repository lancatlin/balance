package balance

import (
	"log"
)

var Num int

func init() {
	Num = 0
}

func sum(numbers, pow []int) int {
	n := 1
	for i, number := range numbers {
		for j := 0; j < pow[i]; j++ {
			n *= number
		}
	}
	if n == -1 {
		log.Fatalln("Number has out of int: ", numbers, pow)
	}
	return n
}

func maximumCommonFactor(a, b int) int {
	Num++
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
