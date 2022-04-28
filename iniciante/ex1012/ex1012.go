package main

import (
	"fmt"
	"math"
)

type valor struct {
	a float64
	b float64
	c float64
}

func areaTriangulo(p valor) {
	areaTriangulo := (p.a * p.c) / 2

	imprimir("TRIANGULO", areaTriangulo)

	areaCirculo(p)
}

func imprimir(forma string, area float64) {
	fmt.Printf("%s: %.3f\n", forma, area)
}

func areaCirculo(p valor) {
	pi := 3.14159
	area := math.Pow(p.c, 2) * pi
	imprimir("CIRCULO:", area)

	areaTrapezio(p)
}

func areaTrapezio(p valor) {
	base := p.a + p.b
	area := (base * p.c) / 2

	imprimir("TRAPEZIO", area)

	areaQuadrado(p)
}

func areaQuadrado(p valor) {
	area := p.b * p.b

	imprimir("QUADRADO", area)

	areaRetangulo(p)
}

func areaRetangulo(p valor) {
	area := p.a * p.b

	imprimir("RETANGULO", area)
}

func main() {
	var valor1 valor

	valor1 = valor{3, 4, 5.2}
	areaTriangulo(valor1)

	valor1 = valor{12.7, 10.4, 15.2}
	areaTriangulo(valor1)
}
