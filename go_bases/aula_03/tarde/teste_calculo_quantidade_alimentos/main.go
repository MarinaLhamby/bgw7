package main

import (
	"fmt"
	"teste_calculo_quantidade_alimentos/food_calculator"
)

func main() {
	animalDog, msg := food_calculator.Animal(food_calculator.Dog)
	if msg != nil {
		panic(msg)
	}
	animalCat, msg := food_calculator.Animal(food_calculator.Cat)
	if msg != nil {
		panic(msg)
	}
	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Println("Quantidade de comida total:", amount, "kg")
}
