package balance

import "log"

func sumCompound(a, b []Compound, ca, cb []int) (aSum, bSum Compound) {
	aSum = make(Compound)
	bSum = make(Compound)
	for i, v := range a {
		log.Println(v, v.multiply(ca[i]))
		aSum.add(v.multiply(ca[i]))
	}
	for i, v := range b {
		bSum.add(v.multiply(cb[i]))
	}
	return
}

func leastCommonMultiple(a, b Compound) Compound {
	for k, v := range b {
		if a[k] < v {
			a[k] = v
		}
	}
	return a
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

func (m Compound) multiply(x int) Compound {
	m1 := make(Compound)
	for k := range m {
		m1[k] = m[k] * x
	}
	return m1
}

func (m Compound) add(m2 Compound) Compound {
	m1 := make(Compound)
	for k, v := range m {
		m1[k] = v
	}
	for k, v := range m2 {
		m1[k] += v
	}
	return m1
}

func (m Compound) total() int {
	n := 0
	for _, v := range m {
		n += v
	}
	return n
}
