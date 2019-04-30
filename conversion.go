package balance

import (
	"log"
	"strconv"
)

func convEquation(equation []rune) (reactants, products []int) {
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
				log.Println(name)
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
