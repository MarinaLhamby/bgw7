package main

import "fmt"

/*
Forma inicial das variáveis:

	var 1firstName string
	var lastName string
	var int age
	1lastName:= 6
	var driver_license = true
	var person height int
	childsNumber:= 2
*/
func main() {
	// Não pode começar com número
	var firstName string
	var lastName string
	// O tipo da variável deve vir após a declaração do nome
	var age int
	// Nome da variável era lastName e não 1lastName e o tipo é string e não int, fora isso não se usa := porque já existe uma variável com esse nome dentro do escopo
	lastName = "Sobrenome"
	// Apesar de não ser um erro o noem driver_license, ele foge do padrão camelCase
	var driverLicense = true
	// Não pode ter espaço no nome da variável
	var personHeight int
	childsNumber := 2

	fmt.Printf("firstName: %s\nlastName: %s\nage: %d\ndriverLicense: %t\npersonHeight: %d\nchildsNumber: %d\n",
		firstName, lastName, age, driverLicense, personHeight, childsNumber)
}
