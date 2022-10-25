package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	DNI          uint64 `json:"dni"`
	FechaIngreso string `json:"fecha"`
}

func (a Alumno) detalle() string {
	return fmt.Sprintf("\nNombre y apellido: %s %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.FechaIngreso)
}

func main() {
	alumno := Alumno{
		Nombre:       "Juan",
		Apellido:     "Calle",
		DNI:          1020460893,
		FechaIngreso: time.Now().UTC().String(),
	}

	fmt.Println(alumno.detalle())

}
