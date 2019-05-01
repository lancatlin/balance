package balance

import (
	"log"
)

func balance(a, b, ca, cb []int) (ra, rb []int, err error) {
	//log.Printf("%v %v\t%v %v\n", a, b, ca, cb)
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

func getInNeeds(a, b uint64) (uint64, uint64) {
	factor := maximumCommonFactor(a, b)
	if factor <= 0 {
		log.Fatalln("multiple is lower than 0:", factor)
	}
	return b / factor, a / factor
}

func selectFactor(a []int, na uint64) int {
	var leastV uint64
	leastV--
	leastI := -1
	for i, v := range a {
		uv := uint64(v)
		l := leastCommonMultiple(na, uv)
		if l < leastV && l != uv*na {
			leastV = l
			leastI = i
		} else if l == na && v > a[leastI] {
			leastI = i
		}
	}
	if leastI == -1 {
		log.Fatalln("Select factor result is -1: ", a, na, leastV)
	}
	return leastI
}
