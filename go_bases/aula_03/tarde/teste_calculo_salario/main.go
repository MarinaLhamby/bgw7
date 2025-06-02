package main

import (
	"fmt"
	"teste_calculo_salario/salary"
)

func main() {
	resultA, err := salary.CalculateSalary(2, salary.A)
	if err != nil {
		fmt.Println("Erro encontrado no caso A:", err)
	}
	resultB, err := salary.CalculateSalary(2, salary.B)
	if err != nil {
		fmt.Println("Erro encontrado no caso B:", err)
	}
	resultC, err := salary.CalculateSalary(2, salary.C)
	if err != nil {
		fmt.Println("Erro encontrado no caso C:", err)
	}
	resultD, err := salary.CalculateSalary(2, "D")
	if err != nil {
		fmt.Println("Erro encontrado no caso D:", err)
	}
	fmt.Println("Resultado A:", resultA)
	fmt.Println("Resultado B:", resultB)
	fmt.Println("Resultado C:", resultC)
	fmt.Println("Resultado D em que ocorre o erro:", resultD)
}
