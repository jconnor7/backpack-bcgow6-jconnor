package main

import "fmt"

func main() {

	var palabra string = "Bootcamp"

	fmt.Printf("Cantidad de letras: %d\n", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%s ", string(palabra[i]))
	}
	fmt.Println()

	for _, value := range palabra {
		fmt.Printf("%s ", string(value))
	}
	fmt.Println()
}
