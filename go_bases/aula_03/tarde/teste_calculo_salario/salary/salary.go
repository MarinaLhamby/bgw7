package salary

import "errors"

const (
	A = "A"
	B = "B"
	C = "C"
)

func CalculateSalary(hours int32, category string) (salary float64, err error) {
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
