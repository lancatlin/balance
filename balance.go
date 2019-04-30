package balance

import (
	"log"
)

var primeNumbers []int

func init() {
	primeNumbers = []int{2, 3, 5, 7, 9, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 51}
}

func balance(a, b, ca, cb []int) (ra, rb []int, err error) {
	log.Println(a, b, ca, cb)
	na, nb := getInNeeds(sum(a, ca), sum(b, cb))
	if na != 1 {
		i := selectFactor(a, na)
		ca[i]++
		return balance(a, b, ca, cb)
	} else if nb != 1 {
		rb, ra, err = balance(b, a, cb, ca)
		return
	} else {
		return ca, cb, nil
	}
}

func getInNeeds(a, b int) (int, int) {
	multiple := leastCommonMultiple(a, b)
	return multiple / a, multiple / b
}

func selectFactor(a []int, na int) int {
	leastV, leastI := int(1e10), -1
	for i, v := range a {
		l := leastCommonMultiple(na, v)
		if l < leastV && l != v*na {
			leastV = l
			leastI = i
		} else if l == na && v > a[leastI] {
			leastI = i
		}
	}
	return leastI
}
