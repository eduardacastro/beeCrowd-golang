package main

import "fmt"

func media(nota1, nota2, nota3 float64) {
	media := (2*nota1 + 3*nota2 + 5*nota3) / 10
	fmt.Printf("MEDIA = %.1f\n", media)
}

func main() {
	media(5, 6, 7)
	media(5, 10, 10)
	media(10, 10, 5)
}
