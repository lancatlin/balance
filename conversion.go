package balance

import (
	"log"
	"strconv"
	"strings"
)

// Compound is the structure of Chemical compound
type Compound map[string]int

func split(equation string) (r, p []string) {
	equation = strings.ReplaceAll(equation, " ", "")
	eq := strings.Split(equation, "=")
	if len(eq) != 2 {
		log.Fatalln("Can't compile this equation: ", equation)
	}
	a, b := eq[0], eq[1]
	r, p = strings.Split(a, "+"), strings.Split(b, "+")
	return
}

func convEquation(equation string) (reactants, products []Compound) {
	r, p := split(equation)
	for _, v := range r {
		reactants = append(reactants, analysis([]rune(v)))
	}
	for _, v := range p {
		products = append(products, analysis([]rune(v)))
	}
	return
}

func export(equation string, ca, cb []int) string {
	r, p := split(equation)
	for i, v := range r {
		if ca[i] > 1 {
			r[i] = strconv.Itoa(ca[i]) + v
		}
	}
	for i, v := range p {
		if cb[i] > 1 {
			p[i] = strconv.Itoa(cb[i]) + v
		}
	}
	reactant := strings.Join(r, " + ")
	products := strings.Join(p, " + ")
	return reactant + " = " + products
}

func sumElements(element, eWeights map[string]int) (sum int) {
	sum = 1
	for key, value := range element {
		for i := 0; i < value; i++ {
			sum *= eWeights[key]
		}
	}
	return
}

func analysis(equation []rune) (elements map[string]int) {
	eh, nh := 0, -1
	elements = make(map[string]int)
	equation = append(equation, 'A')
	for i, v := range equation {
		if v >= 'A' && v <= 'Z' && eh != i {
			if nh == -1 {
				// No number
				name := string(equation[eh:i])
				elements[name]++
			} else {
				name := string(equation[eh:nh])
				value, err := strconv.Atoi(string(equation[nh:i]))
				if err != nil {
					log.Fatal("Conv int fatal: ", string(equation[nh:i]), elements)
				}
				elements[name] += value
			}
			eh = i
			nh = -1
		} else if v >= '0' && v <= '9' && nh == -1 {
			nh = i
		}
	}
	return
}
