package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 1000

	err := checkSalary(salary)
	if err != nil {
		if errors.Is(err, NewSalaryError()) {
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Must pay tax")
}

func checkSalary(salary int) error {
	if salary <= 10000 {
		return NewSalaryError()
	}
	return nil
}

type SalaryError struct {
	Message string
}

func NewSalaryError() SalaryError {
	return SalaryError{
		Message: "salary is less than 10000",
	}
}

func (e SalaryError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}
