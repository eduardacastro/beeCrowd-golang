package main

import "fmt"

type pessoa struct {
	nome    string
	salario float64
	bonus   float64
}

func salario(p pessoa) {
	if p.bonus == 0 {
		imprimir(p.nome, p.salario)
	} else {
		imprimir(p.nome, 0.15*p.bonus+p.salario)
	}
}

func imprimir(nome string, salario float64) {
	fmt.Printf("%s - TOTAL = %.2f\n\n", nome, salario)
}

func main() {
	var pessoa1 pessoa
	pessoa1 = pessoa{
		nome:    "João",
		salario: 500,
		bonus:   1230.30,
	}
	salario(pessoa1)

	var pessoa2 pessoa
	pessoa2 = pessoa{
		nome:    "Pedro",
		salario: 700,
		bonus:   0,
	}
	salario(pessoa2)

	var pessoa3 pessoa
	pessoa3 = pessoa{
		nome:    "MANGOJATA",
		salario: 1700,
		bonus:   1230,
	}
	salario(pessoa3)
}
