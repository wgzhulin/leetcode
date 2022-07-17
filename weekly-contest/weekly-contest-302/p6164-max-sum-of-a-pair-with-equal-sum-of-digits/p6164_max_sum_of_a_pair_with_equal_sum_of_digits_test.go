package max_sum_of_a_pair_with_equal_sum_of_digits

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestMaximumSum(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{Para1: []int{18, 43, 36, 13, 7}},
		{Para1: []int{10, 12, 19, 14}},
		{Para1: []int{383, 77, 97, 261, 102, 344, 150, 130, 55, 337, 401, 498, 21, 5}},
	}

	output := []Output{
		{Para1: 54},
		{Para1: -1},
		{Para1: 460},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, maximumSum(input[i].Para1))
	}
}

func TestBitSum(t *testing.T) {
	assert.Equal(t, 6, bitSum(123))
	assert.Equal(t, 9, bitSum(18))
	assert.Equal(t, 1, bitSum(10))
}
