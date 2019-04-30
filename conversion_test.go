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
