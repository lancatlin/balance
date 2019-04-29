package main

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

func TestMaximumCommonFactor(t *testing.T) {
	type question struct {
		a      int
		b      int
		answer int
	}
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
	type question struct {
		a      int
		b      int
		answer int
	}
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
