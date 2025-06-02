package statistics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperation(t *testing.T) {
	testData := []struct {
		name     string
		op       string
		input    []int64
		expected float64
	}{
		{
			name:     "calculate min value",
			op:       Minimum,
			input:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 1,
		},
		{
			name:     "calculate max value",
			op:       Maximum,
			input:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 9,
		},
		{
			name:     "calculate average value",
			op:       Average,
			input:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 5,
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			function, err := Operation(test.op)

			result := function(test.input...)

			require.Nil(t, err)
			require.Equal(t, test.expected, result)
		})
	}
}
