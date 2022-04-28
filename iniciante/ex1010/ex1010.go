package main

import "fmt"

type peca struct {
	codigo     int
	quantidade int
	valor      float64
}

type compra struct {
	pecas []peca
}

func imprimir(total float64) {
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}

func (p compra) valorTotal() float64 {
	total := 0.0

	for _, peca := range p.pecas {
		total += peca.valor * float64(peca.quantidade)
	}
	return total
}

func main() {
	compra1 := compra{
		pecas: []peca{
			{12, 1, 5.3},
			{16, 2, 5.1},
		},
	}
	imprimir(compra1.valorTotal())

	compra2 := compra{
		pecas: []peca{
			{13, 2, 15.3},
			{161, 4, 5.2},
		},
	}
	imprimir(compra2.valorTotal())

	compra3 := compra{
		pecas: []peca{
			{1, 1, 15.1},
			{2, 1, 15.1},
		},
	}
	imprimir(compra3.valorTotal())
}
