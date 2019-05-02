package balance

var primeNumbers []int

func init() {
	primeNumbers = []int{2, 3, 5, 7, 9, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 51, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
}

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
	if c < d {
		c, d = d, c
	}
	for {
		if d == 0 {
			return c
		}
		c %= d
		c, d = d, c
	}
}

func mcf(a, b, ca, cb []int) (uint64, uint64) {
	aSum := make([]int, len(primeNumbers))
	bSum := make([]int, len(primeNumbers))
	for i, v := range a {
		primeFactorization(v, ca[i], &aSum)
	}
	for i, v := range b {
		primeFactorization(v, cb[i], &bSum)
	}
	na := make([]int, len(primeNumbers))
	nb := make([]int, len(primeNumbers))
	for i := range primeNumbers {
		higher := 0
		if aSum[i] > bSum[i] {
			higher = aSum[i]
		} else {
			higher = bSum[i]
		}
		na[i] = higher - aSum[i]
		nb[i] = higher - bSum[i]
	}
	return sum(primeNumbers, na), sum(primeNumbers, nb)
}

func leastCommonMultiple(aNum, bNum uint64) uint64 {
	factor := maximumCommonFactor(aNum, bNum)
	return aNum * bNum / factor
}

func primeFactorization(number int, x int, factors *[]int) {
	i := 0
	for number != 1 {
		prime := primeNumbers[i]
		if number%prime == 0 {
			(*factors)[i] += x
			number /= prime
		} else {
			i++
		}
	}
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

func sumMap(m1, m2 map[string]int) (mSum map[string]int) {
	mSum = make(map[string]int)
	for key, value := range m1 {
		mSum[key] += value
	}
	for key, value := range m2 {
		mSum[key] += value
	}
	return
}
