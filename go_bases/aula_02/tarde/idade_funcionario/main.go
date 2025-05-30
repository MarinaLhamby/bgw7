package main

import "fmt"

type employee struct {
	Name string
	Age  int
}

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	selectedEmployee := "Benjamin"
	employeeToRemove := "Pedro"
	employeeToAdd := employee{
		Name: "Federico",
		Age:  25,
	}
	age := 21

	fmt.Printf("Idade do %s: %d\n", selectedEmployee, employees[selectedEmployee])
	fmt.Printf("Números de funcionários com idade superior a %d: %d\n", age, countEmployeesAboveAge(employees, age))
	fmt.Printf("Lista antes da adição do %s\n", employeeToAdd.Name)
	printEmployees(employees)
	addEmployee(employees, employeeToAdd)
	fmt.Printf("Lista depois da adição do %s\n", employeeToAdd.Name)
	printEmployees(employees)
	removeEmployee(employees, employeeToRemove)
	fmt.Printf("Lista depois da remoção do %s\n", employeeToRemove)
	printEmployees(employees)
}

func removeEmployee(employees map[string]int, name string) {
	delete(employees, name)
}

func addEmployee(employees map[string]int, emp employee) {
	employees[emp.Name] = emp.Age
}

func countEmployeesAboveAge(employees map[string]int, age int) int {
	sum := 0
	for _, v := range employees {
		if v > age {
			sum++
		}
	}
	return sum
}

func printEmployees(employees map[string]int) {
	for k, v := range employees {
		fmt.Printf("Nome: %s, idade: %d\n", k, v)
	}
}
