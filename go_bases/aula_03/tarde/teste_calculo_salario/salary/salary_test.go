package salary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateSalary(t *testing.T) {
	testData := []struct {
		name     string
		hours    int32
		category string
		expected float64
		err      string
	}{
		{
			name:     "category A salary",
			hours:    120,
			category: A,
			expected: 9000,
		},
		{
			name:     "category B salary",
			hours:    120,
			category: B,
			expected: 3600,
		},
		{
			name:     "category C salary",
			hours:    120,
			category: C,
			expected: 2000,
		},
		{
			name:     "should give error when category its not defined",
			hours:    2,
			category: "D",
			expected: 0,
			err:      "Categoria inválida.",
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			result, err := CalculateSalary(test.hours, test.category)

			require.Equal(t, test.expected, result)
			if err != nil {
				require.Equal(t, test.err, err.Error())
			}
		})
	}
}
