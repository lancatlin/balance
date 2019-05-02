package balance

import (
	"log"
)

func balance(a, b []Compound, ca, cb []int) (ra, rb []int) {
	log.Printf("%v %v\t%v %v\n", a, b, ca, cb)
	na, nb := getInNeeds(a, b, ca, cb)
	if na.total() != 0 {
		i := selectFactor(a, na)
		ca[i]++
		return balance(a, b, ca, cb)
	} else if nb.total() != 0 {
		rb, ra = balance(b, a, cb, ca)
		return
	} else {
		return ca, cb
	}
}

func selectFactor(a []Compound, na Compound) int {
	leastV := int(1e10)
	leastI := -1
	for i, v := range a {
		c := leastCommonMultiple(na, v)
		l := c.total()
		if l < leastV && l != v.add(na).total() {
			leastV = l
			leastI = i
		} else if l == leastV && v.total() > a[leastI].total() {
			leastI = i
		}
	}
	if leastI == -1 {
		log.Fatalln("Select factor result is -1: ", a, na, leastV)
	}
	return leastI
}

func getInNeeds(a, b []Compound, ca, cb []int) (na, nb Compound) {
	aSum, bSum := sumCompound(a, b, ca, cb)
	na, nb = make(Compound), make(Compound)
	for k := range aSum {
		if aSum[k] > bSum[k] {
			nb[k] = aSum[k] - bSum[k]
		} else if aSum[k] < bSum[k] {
			na[k] = bSum[k] - aSum[k]
		}
	}
	return
}
