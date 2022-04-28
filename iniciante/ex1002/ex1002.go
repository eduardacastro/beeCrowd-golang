package main

import (
	"fmt"
	"math"
)

func areaDoCirculo(raio float64) float64 {
	pi := 3.14159
	return math.Pow(raio, 2) * pi
}

func main() {
	//fmt.Printf("%.4f", areaDoCirculo(100.64))
	fmt.Printf("%.4f\n", areaDoCirculo(150))
}
