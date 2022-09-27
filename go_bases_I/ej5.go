package main

import "fmt"

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	fmt.Printf("Edad de Benjamin %d\n", employees["Benjamin"])

	var amountAdult int

	for _, value := range employees {
		if value > 21 {
			amountAdult++
		}
	}
	fmt.Printf("Cantidad de empleados mayores de 21 = %d\n", amountAdult)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
