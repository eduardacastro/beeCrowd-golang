package main

import "fmt"

func imprimir(info string, number float64) {
	if info == "NUMBER" {
		fmt.Printf("\n%s = %.f\n", info, number)
	} else {
		fmt.Printf("%s = U$ %.2f\n", info, number)
	}
}

func salario(horas, valor float64) float64 {
	return horas * valor
}

func main() {
	imprimir("NUMBER", 25)
	imprimir("SALARY", salario(100, 5.50))

	imprimir("NUMBER", 1)
	imprimir("SALARY", salario(200, 20.50))

	imprimir("NUMBER", 6)
	imprimir("SALARY", salario(145, 15.55))
}
