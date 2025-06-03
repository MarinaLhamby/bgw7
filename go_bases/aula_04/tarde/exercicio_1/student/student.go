package student

import (
	"fmt"
	"time"
)

type Student struct {
	Name    string
	Surname string
	DNI     int64
	Date    time.Time
}

func (s Student) Detailing() {
	fmt.Printf("Nome: %s\nSobrenome: %s\nID: %d\nData: %v\n",
		s.Name, s.Surname, s.DNI, s.Date)
}
