package balance

import "log"

func init1Slice(l int) []int {
	slice := make([]int, l, l)
	for i := range slice {
		slice[i] = 1
	}
	return slice
}

// Balance main program
func Balance(equation string) (result string) {
	a, b := convEquation(equation)
	log.Println(a, b)
	ca, cb := balance(a, b, init1Slice(len(a)), init1Slice(len(b)))
	return export(equation, ca, cb)
}
