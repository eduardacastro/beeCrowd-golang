package main

import "fmt"

func multiplicacao(num1, num2 int) int {
	return num1 * num2
}

func diferenca(ab, cd int) int {
	return ab - cd
}

func imprimir(diferenca int) {
	fmt.Printf("DIFERENÇA = %v\n", diferenca)
}

func main() {
	imprimir(diferenca(multiplicacao(5, 6), multiplicacao(7, 8)))
	imprimir(diferenca(multiplicacao(0, 0), multiplicacao(7, 8)))
	imprimir(diferenca(multiplicacao(5, 6), multiplicacao(-7, 8)))
}
