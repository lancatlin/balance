package balance

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	a := Compound{"C": 1, "O": 2}
	b := a.multiply(2)
	if b["C"] != 2 || b["O"] != 4 {
		t.Error("Multiply fatal: ", a, b)
	}
	if a["C"] != 1 || a["O"] != 2 {
		t.Error("A has been change", a)
	}
}