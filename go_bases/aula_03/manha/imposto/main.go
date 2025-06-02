package main

import "fmt"

func main() {
	salary1, salary2, salary3, salary4, salary5 := 1000.00, 50000.00, 60000.00, 150000.00, 160000.00
	calculateTax(salary1)
	calculateTax(salary2)
	calculateTax(salary3)
	calculateTax(salary4)
	calculateTax(salary5)

}
func calculateTax(salary float64) {
	tax := 0.0
	fee := 0.0
	if salary > 150000 {
		fee = 0.27
	} else if salary > 50000 {
		fee = 0.17
	} else {
		fee = 0
	}

	tax = salary * fee
	fmt.Printf("Valor do salário: %.2f, valor do imposto: %.2f, valor da taxa aplicada como imposto: %.2f\n", salary, tax, fee)
}
