package main

import "fmt"

func somar(x, y int) int {
	return x + y
}

func imprimir(soma int) {
	fmt.Println(soma)
}

func main() {
	// x := -10
	// y := 4

	x := 15
	y := -7

	imprimir(somar(x, y))

}
