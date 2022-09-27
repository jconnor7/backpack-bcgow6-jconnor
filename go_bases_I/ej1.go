package main

import "fmt"

func main() {
	nombre, direccion := "Juan Calle", "Cra55#25-99"
	fmt.Printf("Nombre:  %v \n", nombre)
	fmt.Printf("Dirección: %v \n", direccion)
	var (
		temperatura = 24  // °F
		humedad     = 58  // HR
		presion     = 640 // mmHg
	)
	fmt.Printf("Temperatura: %v \n", temperatura)
	fmt.Printf("Presion: %v \n", presion)
	fmt.Printf("Humedad: %v \n", humedad)
}
