package file

import (
	"testing"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/pkg/apperrors"
	"github.com/stretchr/testify/require"
)

const (
	pathToTestFile = "./data/tickets_test.csv"
)

func TestGetTicketsFromFile(t *testing.T) {
	dataTest := []struct {
		name        string
		path        string
		expected    int
		expectedErr error
	}{
		{
			name:     "Successfully read tickets from file",
			path:     pathToTestFile,
			expected: 1000,
		},
		{
			name:        "File not found",
			path:        "../tickets_not_found.csv",
			expected:    0,
			expectedErr: &apperrors.ErrReadingFile,
		},
	}
	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetTicketsFromFile(tt.path)
			require.Equal(t, tt.expected, len(result))
			if err != nil {
				require.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}
