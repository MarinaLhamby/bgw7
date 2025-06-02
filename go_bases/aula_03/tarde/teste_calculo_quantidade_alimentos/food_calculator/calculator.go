package food_calculator

import "errors"

const (
	Dog       = "Dog"
	Cat       = "Cat"
	Hamster   = "Hamster"
	Tarantula = "Tarantula"
)

var foodPerAnimalInKg = map[string]float64{
	Dog:       10,
	Cat:       5,
	Hamster:   0.25,
	Tarantula: 0.15,
}

func Animal(animal string) (func(numberofAnimals int) float64, error) {
	switch animal {
	case Dog:
		return CalculateDogFood, nil
	case Cat:
		return CalculateCatFood, nil
	case Tarantula:
		return CalculateTarantulaFood, nil
	case Hamster:
		return CalculateHamsterFood, nil
	default:
		return nil, errors.New("Animal inválido.")
	}
}

func CalculateDogFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[Dog] * float64(numberOfAnimals)
}

func CalculateCatFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[Cat] * float64(numberOfAnimals)
}

func CalculateTarantulaFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[Tarantula] * float64(numberOfAnimals)
}

func CalculateHamsterFood(numberOfAnimals int) float64 {
	return foodPerAnimalInKg[Hamster] * float64(numberOfAnimals)
}
