package balance 

import (
	"testing"
)

func TestSum(t *testing.T) {
	q := []int{3, 5, 7}
	c := []int{2, 1, 3}
	answers := sum(q, c)
	if answers != 15435 {
		t.Error("sum answer is not correct")
	}
}

type question struct {
	a      uint64
	b      uint64
	answer uint64
}

func TestMaximumCommonFactor(t *testing.T) {
	questions := []question{
		question{12, 18, 6},
		question{24, 8, 8},
		question{64, 144, 16},
	}
	for _, v := range questions {
		if a := maximumCommonFactor(v.a, v.b); a != v.answer {
			t.Errorf("Maximum common factor is not correct: %d, %d: answer: %d\t get: %d", v.a, v.b, v.answer, a)
		}
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	questions := []question{
		question{12, 18, 36},
		question{24, 8, 24},
		question{64, 144, 576},
	}
	for _, v := range questions {
		if a := leastCommonMultiple(v.a, v.b); a != v.answer {
			t.Errorf("Least common multiple is not correct: %d, %d: answer: %d\t get: %d", v.a, v.b, v.answer, a)
		}
	}
}

func TestSumMap(t *testing.T) {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	m1["C"] = 1
	m1["O"] = 2
	m2["H"] = 2
	m2["O"] = 1
	mSum := sumMap(m1, m2)
	if !(mSum["C"] ==1 && mSum["O"] == 3 && mSum["H"] == 2) {
		t.Error("Sum map fatal: ", m1, m2, mSum)
	}
}