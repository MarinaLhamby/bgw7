package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		panic(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Resultados:\nminValue:%.2f\naverageValue:%.2f\nmaxValue:%.2f\n", minValue, averageValue, maxValue)
}

func operation(op string) (func(values ...int) float64, error) {
	switch op {
	case minimum:
		return opMinimum, nil
	case average:
		return opAverage, nil
	case maximum:
		return opMaximum, nil
	default:
		return nil, errors.New("Função não definida.")
	}
}

func opMinimum(values ...int) float64 {
	smallest := values[0]
	for _, v := range values {
		if v < smallest {
			smallest = v
		}
	}
	return float64(smallest)
}

func opAverage(values ...int) float64 {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return float64(sum) / float64(len(values))
}

func opMaximum(values ...int) float64 {
	highest := values[0]
	for _, v := range values {
		if v > highest {
			highest = v
		}
	}
	return float64(highest)
}
