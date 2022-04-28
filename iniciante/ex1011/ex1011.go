package main

import (
	"fmt"
	"math"
)

func volumeDaEsfera(raio float64) {
	pi := 3.14159
	variavel1 := 4 * pi
	variavel := variavel1 / 3

	imprimir(math.Pow(raio, 3) * variavel)
}

func imprimir(volume float64) {
	fmt.Printf("VOLUME = %.3f\n", volume)
}

func main() {

	volumeDaEsfera(3)
	volumeDaEsfera(15)
	volumeDaEsfera(1523)
}
