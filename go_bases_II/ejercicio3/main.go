package main

import (
	"fmt"
)

const (
	catA = "A"
	catB = "B"
	catC = "C"
)

func salario(mins float64, cat string) (float64, string) {
	hrs := mins / 60
	switch cat {
	case catA:
		return hrs * 3000 * 1.5, ""
	case catB:
		return hrs * 1500 * 1.2, ""
	case catC:
		return hrs * 1000, ""
	default:
		return 0, "Los datos son incorrectos"
	}
}

func main() {
	var minsPerWorkDay float64 = 8 * 60 // 8 hrs * 60 mins
	var category string = "B"

	salary, msg := salario(minsPerWorkDay, category)
	if msg != "" {
		fmt.Println(msg)
	} else {
		fmt.Printf("Categoria: %s, Sueldo: %0.2f\n", category, salary)
	}
}
