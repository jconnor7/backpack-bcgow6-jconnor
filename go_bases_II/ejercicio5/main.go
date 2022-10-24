package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarántula"
)

func animalDog(dogs int) float64 {
	return 10 * float64(dogs)
}

func animalCat(cats int) float64 {
	return 5 * float64(cats)
}

func animalTarantula(tarantula int) float64 {
	return (150 * float64(tarantula)) / 1000
}

func animalHamster(hamster int) float64 {
	return (250 * float64(hamster)) / 1000
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		msgError := fmt.Sprintf("El animal (%s) no esta registrado", animal)
		return nil, errors.New(msgError)
	}
}

func main() {
	var cantidad float64

	animal, err := Animal(dog)
	numDogs := 2
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animal(numDogs)
	}

	animalcat, err := Animal(cat)
	numCats := 5
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalcat(numCats)
	}

	animalhamster, err := Animal(hamster)
	numHamsters := 4
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalhamster(numHamsters)
	}

	animaltarantula, err := Animal(tarantula)
	numTarantula := 3
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animaltarantula(numTarantula)
	}
	fmt.Println("Cantidad de animales")
	fmt.Println("Perros:", numDogs)
	fmt.Println("Gatos:", numCats)
	fmt.Println("Hamsters:", numHamsters)
	fmt.Println("Tarantulas:", numTarantula)
	fmt.Println("Total de alimentos necesarios", cantidad, "Kg")
}
