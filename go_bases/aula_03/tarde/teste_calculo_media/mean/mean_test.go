package mean

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateMean(t *testing.T) {
	testData := []struct {
		name     string
		input    []float32
		expected float32
	}{
		{
			name:     "calcular média dos alunos",
			input:    []float32{10, 2, 6, 8, 9},
			expected: 7.0,
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateMean(test.input...)

			require.Equal(t, test.expected, result)
		})
	}
}
