package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func opMinimum(valores ...float64) float64 {
	var min float64 = valores[0]

	for _, valor := range valores {

		if valor < min {
			min = valor
		}
	}

	return min
}

func opMaximum(valores ...float64) float64 {
	var max float64 = valores[0]

	for _, valor := range valores {

		if valor > max {
			max = valor
		}
	}

	return max
}

func opAverage(valores ...float64) float64 {
	var promedio float64
	var totalNotas float64

	for _, valor := range valores {
		totalNotas += valor
	}

	promedio = totalNotas / float64(len(valores))
	return promedio
}

func operation(function string) (func(...float64) float64, error) {
	switch function {
	case minimum:
		return opMinimum, nil
	case maximum:
		return opMaximum, nil
	case average:
		return opAverage, nil
	default:
		msgError := fmt.Sprintf("La función (%s) no ha sido definida.", function)
		return nil, errors.New(msgError)
	}
}

func main() {
	minValue, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}

	averageValue, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxValue, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Promedio %.2f\n", averageValue(10, 7, 10, 10)) // 9.25
	fmt.Printf("Minimo %.2f\n", minValue(1, 10, 10, 4))        // 1.00
	fmt.Printf("Maximo %.2f\n", maxValue(10, 3, 10, 1))        // 10.00
}
