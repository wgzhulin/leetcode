package make_array_zero_by_subtracting_equal_amounts

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestMinimumOperations(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{Para1: []int{1, 5, 0, 3, 5}},
		{Para1: []int{0}},
	}

	output := []Output{
		{Para1: 3},
		{Para1: 0},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, minimumOperations(input[i].Para1))
	}
}
