package main

import (
	"fmt"
)

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
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}
	return nil
}
