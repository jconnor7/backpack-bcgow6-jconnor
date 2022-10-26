package main

import "fmt"

type Matrix struct {
	values []float64
	height int
	width  int
}

func (m Matrix) Set() {
	if len(m.values) != m.height*m.width {
		fmt.Printf("Verificar datos de matriz no coinciden con la dimension")
	}
}

func (m Matrix) IsSquare() bool {
	if m.width == m.height && m.height != 0 && m.width != 0 {
		return true
	}
	return false
}

func (m Matrix) MaxValue() float64 {
	max := m.values[0]

	for _, value := range m.values {
		if value > max {
			max = value
		}
	}
	return max
}

func (m Matrix) Print() {
	if len(m.values) == 0 {
		fmt.Printf("Matriz vacia")
	}
	for row := 0; row < m.height; row++ {
		startRow := row * m.width
		endRow := startRow + m.width
		fmt.Printf("%f \n", m.values[startRow:endRow])
	}

}

func main() {

	m := Matrix{
		values: []float64{0, 5, 3, 7, 13, 11, 9, 8, 4},
		height: 3,
		width:  3,
	}

	Matrix.Set(m)
	Matrix.Print(m)
	Matrix.IsSquare(m)
}
