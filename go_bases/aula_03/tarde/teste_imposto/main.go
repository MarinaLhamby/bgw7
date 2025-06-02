package main

import (
	"fmt"
	tax "teste_imposto/tax"
)

func main() {
	salary1, salary2, salary3, salary4, salary5 := 1000.00, 50000.00, 60000.00, 150000.00, 160000.00
	fmt.Printf("Valor do salário: %.2f, valor da taxa aplicada como imposto: %.2f\n", salary1, tax.CalculateTax(salary1))
	fmt.Printf("Valor do salário: %.2f, valor da taxa aplicada como imposto: %.2f\n", salary2, tax.CalculateTax(salary2))
	fmt.Printf("Valor do salário: %.2f, valor da taxa aplicada como imposto: %.2f\n", salary3, tax.CalculateTax(salary3))
	fmt.Printf("Valor do salário: %.2f, valor da taxa aplicada como imposto: %.2f\n", salary4, tax.CalculateTax(salary4))
	fmt.Printf("Valor do salário: %.2f, valor da taxa aplicada como imposto: %.2f\n", salary5, tax.CalculateTax(salary5))
}
