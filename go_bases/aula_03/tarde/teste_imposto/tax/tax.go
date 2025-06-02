package tax

func CalculateTax(salary float64) float64 {
	tax := 0.0
	fee := 0.0
	if salary > 150000 {
		fee = 0.27
	} else if salary > 50000 {
		fee = 0.17
	} else {
		fee = 0
	}

	tax = salary * fee
	return tax
}
