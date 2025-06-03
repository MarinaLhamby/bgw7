package main

import (
	"exercicio_2/employee"
	"time"
)

func main() {
	person := employee.Person{
		ID:          1,
		Name:        "Marina",
		DateOfBirth: time.Date(1997, 10, 26, 0, 0, 0, 0, &time.Location{}),
	}
	employee := employee.Employee{
		ID:       1,
		Position: "Software Developer",
		Person:   person,
	}

	employee.PrintEmployee()
}
