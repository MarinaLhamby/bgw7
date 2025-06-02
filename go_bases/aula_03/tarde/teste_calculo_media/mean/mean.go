package mean

func CalculateMean(grades ...float32) float32 {
	numberOfStudents := len(grades)
	var sum float32 = 0.0
	for _, g := range grades {
		sum += g
	}

	return sum / float32(numberOfStudents)
}
