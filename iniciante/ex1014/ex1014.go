package main

import (
	"fmt"
)

func main() {
	// distancia := 500
	// gasolina := 35.0
	distancia := 2254
	gasolina := 124.4

	divisao := float64(distancia) / gasolina
	fmt.Printf("%.3f km/l\n", divisao)
}
