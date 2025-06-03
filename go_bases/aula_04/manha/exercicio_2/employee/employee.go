package employee

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int64
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID       int64
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("%+v", *e)
}
