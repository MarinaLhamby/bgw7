package main

import (
	"fmt"
	"teste_calculo_estatistica/statistics"
)

func main() {
	minFunc, err := statistics.Operation(statistics.Minimum)
	if err != nil {
		panic(err)
	}
	averageFunc, err := statistics.Operation(statistics.Average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := statistics.Operation(statistics.Maximum)
	if err != nil {
		panic(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Resultados:\nminValue:%.2f\naverageValue:%.2f\nmaxValue:%.2f\n", minValue, averageValue, maxValue)
}
