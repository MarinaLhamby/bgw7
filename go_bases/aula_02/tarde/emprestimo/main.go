package main

import "fmt"

type entrada struct {
	Idade        int
	Empregado    bool
	TempoEmprego int
	Salario      float64
}

func main() {
	entradas := []entrada{
		{
			Idade:        23,
			Empregado:    true,
			TempoEmprego: 2,
			Salario:      200000,
		},
		{
			Idade:        23,
			Empregado:    true,
			TempoEmprego: 2,
			Salario:      100000,
		},
		{
			Idade:        23,
			Empregado:    true,
			TempoEmprego: 1,
			Salario:      200000,
		},
		{
			Idade:     23,
			Empregado: false,
		},
		{
			Idade: 12,
		},
	}

	for _, e := range entradas {
		avaliarEmprestimo(e)
	}
}

/*
Regras:

	>22 anos
	empregados
	empregados há > 1 ano
	Caso aprovado:
		Juros não são cobrados quando salário > 100.000
*/
func avaliarEmprestimo(input entrada) {
	switch {
	case input.Idade > 22 && input.Empregado && input.TempoEmprego > 1:
		if input.Salario > 100000 {
			fmt.Println("Empréstimo aprovado sem juros.")
		} else {
			fmt.Println("Empréstimo aprovado.")
		}
	case input.Idade <= 22:
		fmt.Println("Empréstimo reprovado por idade.")
	case !input.Empregado:
		fmt.Println("Empréstimo reprovado por situação empregatícia.")
	case input.Empregado && input.TempoEmprego <= 1:
		fmt.Println("Empréstimo reprovado por tempo de emprego.")
	default:
		fmt.Println("Empréstimo reprovado.")
	}
}
