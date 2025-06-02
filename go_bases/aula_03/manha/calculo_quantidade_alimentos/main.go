package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

var foodPerAnimalInKg = map[string]float64{
	dog:       10,
	cat:       5,
	hamster:   0.25,
	tarantula: 0.15,
}

func main() {
	animalDog, msg := animal(dog)
	if msg != nil {
		panic(msg)
	}
	animalCat, msg := animal(cat)
	if msg != nil {
		panic(msg)
	}
	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Println("Quantidade de comida total:", amount, "kg")
}

func animal(animal string) (func(numberofAnimals int) float64, error) {
	switch animal {
	case dog:
		return calculateDogFood, nil
	case cat:
		return calculateCatFood, nil
	case tarantula:
		return calculateTarantulaFood, nil
	case hamster:
		return calculateHamsterFood, nil
	default:
		return nil, errors.New("Animal inválido.")
	}
}

func calculateDogFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[dog] * float64(numberOfAnimals)
}

func calculateCatFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[cat] * float64(numberOfAnimals)
}

func calculateTarantulaFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[tarantula] * float64(numberOfAnimals)
}

func calculateHamsterFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[hamster] * float64(numberOfAnimals)
}
