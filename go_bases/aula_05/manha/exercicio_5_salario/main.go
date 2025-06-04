package main

import (
	"errors"
	"fmt"
)

var InvalidWorkHours = errors.New(
	"Error: the worker cannot have worked less than 80 hours per month")

func main() {
	salary := 100000.0
	hours := 70.0

	calculatedSalary, err := calculateSalary(hours, salary)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("O salário calculado foi de: %.2f\n", calculatedSalary)
}

func calculateSalary(
	hours float64, salaryPerHours float64) (calculatedSalary float64, err error) {
	if hours < 80 {
		return 0, InvalidWorkHours
	}

	calculatedSalary = salaryPerHours * hours

	if calculatedSalary >= 150000 {
		calculatedSalary -= calculatedSalary * 0.1
	}

	return calculatedSalary, nil
}
