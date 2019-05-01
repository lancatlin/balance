package balance

import (
	"testing"
)

func TestAnalysis(t *testing.T) {
	type testData struct {
		q []rune
		a map[string]int
	}
	data := []testData{
		testData{
			[]rune("C6H12O6"),
			map[string]int{"C": 6, "H": 12, "O": 6},
		},
		testData{
			[]rune("Fe2O3"),
			map[string]int{"Fe": 2, "O": 3},
		},
		testData{
			[]rune("C2H5OH"),
			map[string]int{"C": 2, "H": 6, "O": 1},
		},
	}
	for _, v := range data {
		answer := analysis(v.q)
		for name, number := range v.a {
			if value, ok := answer[name]; ok {
				if value != number {
					t.Errorf("Answer is not equal: should %s: %d\tget %s: %d", name, number, name, value)
					t.Log(answer)
				}
			} else {
				t.Errorf("Answer is not equal: doesn't find %s", name)
				t.Log(answer)
			}
		}
	}
}

func TestConvEquation(t *testing.T) {
	q := "C2H5OH + O2 = CO2 + H2O"
	reactants, products := convEquation(q)
	if !equal(reactants, []int{4800, 9}) || !equal(products, []int{45, 12}) {
		t.Error("Conversion Fatal: ", reactants, products)
	}
}

func TestExport(t *testing.T) {
	equation := "C2H5OH + O2 = CO2 + H2O"
	ca, cb := []int{1, 3}, []int{2, 3}
	if a := export(equation, ca, cb); a != "C2H5OH + 3O2 = 2CO2 + 3H2O" {
		t.Error("Export fatal: ", a)
	}
}
