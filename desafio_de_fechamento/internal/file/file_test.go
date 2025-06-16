package file

import (
	"testing"
	"time"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/internal_error"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTicketsFromFile(t *testing.T) {
	time, _ := time.Parse("15:04", "17:11")
	dataTest := []struct {
		name        string
		path        string
		expected    []tickets.Ticket
		expectedErr error
	}{
		{
			name: "Successfully read tickets from file",
			path: "../../tickets_test.csv",
			expected: []tickets.Ticket{
				{
					ID:            1,
					Name:          "Tait Mc Caughan",
					Email:         "tmc0@scribd.com",
					Destination:   "Finland",
					DepartureTime: time,
					Price:         785,
				},
				{
					ID:            2,
					Name:          "Tait Mc Caughan",
					Email:         "tmc0@scribd.com",
					Destination:   "Finland",
					DepartureTime: time,
					Price:         785,
				},
			},
		},
		{
			name:        "File not found",
			path:        "../tickets_not_found.csv",
			expected:    nil,
			expectedErr: &internal_error.ErrReadingFile,
		},
	}
	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetTicketsFromFile(tt.path)
			require.Equal(t, tt.expected, result)
			if err != nil {
				require.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}
