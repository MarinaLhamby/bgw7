package statistics

import "errors"

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func Operation(op string) (func(values ...int64) float64, error) {
	switch op {
	case Minimum:
		return opMinimum, nil
	case Average:
		return opAverage, nil
	case Maximum:
		return opMaximum, nil
	default:
		return nil, errors.New("Função não definida.")
	}
}

func opMinimum(values ...int64) float64 {
	smallest := values[0]
	for _, v := range values {
		if v < smallest {
			smallest = v
		}
	}
	return float64(smallest)
}

func opAverage(values ...int64) float64 {
	var sum int64 = 0
	for _, v := range values {
		sum += v
	}
	return float64(sum) / float64(len(values))
}

func opMaximum(values ...int64) float64 {
	highest := values[0]
	for _, v := range values {
		if v > highest {
			highest = v
		}
	}
	return float64(highest)
}
