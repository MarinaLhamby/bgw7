package main

import (
	"errors"
	"fmt"
)

const (
	A = "A"
	B = "B"
	C = "C"
)

func main() {
	resultA, err := calculateSalary(2, A)
	if err != nil {
		fmt.Println("Erro encontrado no caso A:", err)
	}
	resultB, err := calculateSalary(2, B)
	if err != nil {
		fmt.Println("Erro encontrado no caso B:", err)
	}
	resultC, err := calculateSalary(2, C)
	if err != nil {
		fmt.Println("Erro encontrado no caso C:", err)
	}
	resultD, err := calculateSalary(2, "D")
	if err != nil {
		fmt.Println("Erro encontrado no caso D:", err)
	}
	fmt.Println("Resultado A:", resultA)
	fmt.Println("Resultado B:", resultB)
	fmt.Println("Resultado C:", resultC)
	fmt.Println("Resultado D em que ocorre o erro:", resultD)
}

func calculateSalary(hours int32, category string) (salary float64, err error) {
	hFloat := float64(hours)
	switch category {
	case A:
		salary = 3000 * hFloat
		salary += salary * 0.5
	case B:
		salary = 1500 * hFloat
		salary += salary * 0.2
	case C:
		salary = 1000 * hFloat
	default:
		err = errors.New("Categoria inválida.")
	}
	return
}
