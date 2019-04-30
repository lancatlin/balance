package balance

import (
	"testing"
)

func init1Slice(l int) []int {
	slice := make([]int, l, l)
	for i := range slice {
		slice[i] = 1
	}
	return slice
}

func TestBalance(t *testing.T) {
	type Equation struct {
		Qr []int
		Qp []int
		Ar []int
		Ap []int
	}
	testData := []Equation{
		// "C2H5OH + 3O2 → 2CO2 + 3H2O"
		Equation{
			[]int{4800, 9},
			[]int{45, 12},
			[]int{1, 3},
			[]int{2, 3},
		},
		// "CH4 + 2O2 → CO2 + 2H2O"
		Equation{
			[]int{80, 9},
			[]int{45, 12},
			[]int{1, 2},
			[]int{1, 2},
		},
		// "C3H8 + 5O2 → 3CO2 + 4H2O"
		Equation{
			[]int{32000, 9},
			[]int{45, 12},
			[]int{1, 5},
			[]int{3, 4},
		},
		// "2Fe2O3 + 3C → 4Fe + 3CO2"
		Equation{
			[]int{200, 3},
			[]int{5, 12},
			[]int{2, 3},
			[]int{4, 3},
		},
	}
	for _, v := range testData {
		ra, rb, err := balance(v.Qr, v.Qp, init1Slice(len(v.Qr)), init1Slice(len(v.Qp)))
		if err != nil {
			t.Log(err)
		}
		if !equal(ra, v.Ar) || !equal(rb, v.Ap) {
			t.Errorf("Answer is not correct: ra: %d\trb: %d\t\n", ra, rb)
		}
	}
}
