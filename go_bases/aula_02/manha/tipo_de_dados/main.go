package main

import "fmt"

/*
Tipos de variáveis iniciais:

	var lastName string = "Smith"
	var age int = "35"
	var salary string = 45857.90
	var firstName string = "Marina"
*/
func main() {
	var lastName string = "Smith"
	// Declarado como int mas estava sendo atribuído uma string
	var age int = 35
	// Declarado como string mas atribuído float
	var salary float64 = 45857.90
	var firstName string = "Mary"

	fmt.Printf("lastName: %s\nage: %d\nsalary: %.2f\nfirstName: %s\n", lastName, age, salary, firstName)
}
