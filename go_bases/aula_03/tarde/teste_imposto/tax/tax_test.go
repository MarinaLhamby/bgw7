package tax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTax(t *testing.T) {
	testData := []struct {
		name     string
		input    float64
		expected float64
	}{
		{name: "salary below 50.000", input: 40000.0, expected: 0.0},
		{name: "salary above 50.000", input: 60000.0, expected: 10200.0},
		{name: "salary above 150.000", input: 151000.0, expected: 40770.0},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateTax(test.input)

			require.Equal(t, test.expected, result)
		})
	}
}
