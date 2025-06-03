package main

import (
	"exercicio_1/student"
	"time"
)

func main() {
	students := []student.Student{
		{
			Name:    "Student 1",
			Surname: "Surname 1",
			DNI:     1,
			Date:    time.Now(),
		},
		{
			Name:    "Student 2",
			Surname: "Surname 2",
			DNI:     2,
			Date:    time.Now(),
		},
	}

	for _, student := range students {
		student.Detailing()
	}
}
