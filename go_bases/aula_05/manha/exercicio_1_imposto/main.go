package main

import "fmt"

func main() {
	salary := 1000
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return NewSalaryError()
	}
	return nil
}

type SalaryError struct {
	Message string
}

func NewSalaryError() *SalaryError {
	return &SalaryError{
		Message: "the salary entered does not reach the taxable minimum",
	}
}

func (me SalaryError) Error() string {
	return fmt.Sprintf("Error: %s", me.Message)
}
