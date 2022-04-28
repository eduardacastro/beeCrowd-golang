package main

import "fmt"

func media(nota1, nota2 float64) {
	media := (3.5*nota1 + 7.5*nota2) / 11
	fmt.Printf("MEDIA = %.5f\n", media)
}

func main() {
	media(5, 7.1)
	media(0, 7.1)
	media(10, 10)
}
