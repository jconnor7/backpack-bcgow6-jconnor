package main

import "fmt"

func calculateTaxes(salary float32) float32 {
	switch {
	case salary > 150000:
		{
			return salary * 0.27
		}
	case salary > 50000:
		{
			return salary * 0.17
		}
	default:
		return 0
	}
}

func main() {

	var salary float32 = 150000

	fmt.Printf("Descontar impuesto a %.2f por un valor igual: %.2f\n", salary, calculateTaxes(salary))
}
