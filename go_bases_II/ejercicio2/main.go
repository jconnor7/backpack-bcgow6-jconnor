package main

import "fmt"

func meanGrades(grades ...float32) (float32, string) {
	var sumGrades float32

	for i, grade := range grades {
		if grade < 0 {
			return 0, fmt.Sprintf("Nota %d con valor %.2f es inválida", i+1, grade)
		}
		sumGrades += grade
	}

	return sumGrades / float32(len(grades)), ""
}

func main() {
	mean, msg := meanGrades(9, 9, 7, 9, 8)

	if msg != "" {
		fmt.Printf("%s\n", msg)
		return
	}

	fmt.Printf("El promedio de calificaciones es: %.2f\n", mean)
}
