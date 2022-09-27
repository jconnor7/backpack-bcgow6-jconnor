package main

import "fmt"

func main() {
	// Rules:
	var (
		vAge       int     = 22
		vEmployee  bool    = true
		vSeniority int     = 1
		vSalary    float64 = 100000.00
	)

	// Data Employee:
	var (
		age       = 29
		employee  = true
		seniority = 3
		salary    = 500000.00
	)

	switch {
	case age <= vAge:
		fmt.Println("Debes ser mayor de 22 años de edad.")
	case employee != vEmployee:
		fmt.Println("Debes estar empleado.")
	case seniority < vSeniority:
		fmt.Println("Debes tener una antiguedad mayor a 1 año")
	case salary < vSalary:
		fmt.Println("Otorgar credito sin interés ✅")
	default:
		fmt.Println("Otorgar credito con interés ✔️")
	}

}
