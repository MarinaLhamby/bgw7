package food_calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateDogFood(t *testing.T) {
	numberOfDogs := 8
	expected := 80.0

	result := CalculateDogFood(numberOfDogs)

	require.Equal(t, expected, result)
}

func TestCalculateCatFood(t *testing.T) {
	numberOfCats := 2
	expected := 10.0

	result := CalculateCatFood(numberOfCats)

	require.Equal(t, expected, result)
}

func TestCalculateHamsterFood(t *testing.T) {
	numberOfHamsters := 3
	expected := 0.75

	result := CalculateHamsterFood(numberOfHamsters)

	require.Equal(t, expected, result)
}

func TestCalculateTarantulaFood(t *testing.T) {
	numberOfTarantulas := 3
	expected := 0.45

	result := CalculateTarantulaFood(numberOfTarantulas)

	require.InDelta(t, expected, result, 1e-6)
}
