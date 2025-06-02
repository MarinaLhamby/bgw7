package main

import "fmt"

func main() {
	fmt.Printf("Média das notas dos alunos: %.2f\n", calculateMean(10, 8, 2, 3, 6))
	fmt.Printf("Média das notas dos alunos: %.2f\n", calculateMean(10, 10, 10, 10, 10))
	fmt.Printf("Média das notas dos alunos: %.2f\n", calculateMean(10, 8, 6, 3, 6))

}

func calculateMean(grades ...float32) float32 {
	numberOfStudents := len(grades)
	var sum float32 = 0.0
	for _, g := range grades {
		sum += g
	}

	return sum / float32(numberOfStudents)
}
