package main

import (
	"fmt"
	mean "teste_calculo_media/mean"
)

func main() {
	fmt.Printf("Média das notas dos alunos: %.2f\n", mean.CalculateMean(10, 8, 2, 3, 6))
	fmt.Printf("Média das notas dos alunos: %.2f\n", mean.CalculateMean(10, 10, 10, 10, 10))
	fmt.Printf("Média das notas dos alunos: %.2f\n", mean.CalculateMean(10, 8, 6, 3, 6))
}
