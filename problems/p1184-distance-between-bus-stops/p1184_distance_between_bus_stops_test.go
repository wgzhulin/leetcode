package distance_between_bus_stops

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistanceBetweenBusStops(t *testing.T) {
	type Input struct {
		Para1 []int
		Para2 int
		Para3 int
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{
			Para1: []int{1, 2, 3, 4},
			Para2: 0,
			Para3: 1,
		},
		{
			Para1: []int{1, 2, 3, 4},
			Para2: 0,
			Para3: 2,
		},
		{
			Para1: []int{1, 2, 3, 4},
			Para2: 0,
			Para3: 3,
		},
		{
			Para1: []int{7, 10, 1, 12, 11, 14, 5, 0},
			Para2: 7,
			Para3: 2,
		},
	}

	output := []Output{
		{Para1: 1},
		{Para1: 3},
		{Para1: 4},
		{Para1: 17},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, distanceBetweenBusStops(input[i].Para1,
			input[i].Para2, input[i].Para3))
	}
}
