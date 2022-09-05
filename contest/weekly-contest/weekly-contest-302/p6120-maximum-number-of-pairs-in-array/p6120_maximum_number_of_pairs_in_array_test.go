package maximum_number_of_pairs_in_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfPairs(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 []int
	}

	input := []Input{
		{
			Para1: []int{1, 3, 2, 1, 3, 2, 2},
		},
		{
			Para1: []int{1, 1},
		},
		{
			Para1: []int{0},
		},
	}

	output := []Output{
		{Para1: []int{3, 1}},
		{Para1: []int{1, 0}},
		{Para1: []int{0, 1}},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, numberOfPairs(input[i].Para1))
	}
}
