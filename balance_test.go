package main

import (
	"testing"
)

type Equation struct {
	Qr []int
	Qp []int
	Ar []int
	Ap []int
}

func TestBalance(t *testing.T) {
	testData := []Equation{
		Equation{
			[]int{4800, 9},
			[]int{15, 12},
			[]int{1, 3},
			[]int{2, 3},
		},
	}
	t.Log(testData[0])
}
