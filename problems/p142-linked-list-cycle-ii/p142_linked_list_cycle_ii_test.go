package linked_list_cycle_ii

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	type Input struct {
		Para1 []int
		Pos   int
	}

	type Output struct {
		Para1 []int
	}

	input := []Input{
		{
			Para1: []int{3, 2, 0, -4},
			Pos:   1,
		},
		{
			Para1: []int{1, 2},
			Pos:   0,
		},
		{
			Para1: []int{1},
			Pos:   -1,
		},
		{
			Para1: []int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5},
			Pos:   6,
		},
	}

	output := []Output{
		{Para1: []int{2, 0, -4}},
		{Para1: []int{1, 2}},
		{Para1: []int{}},
		{Para1: []int{-9, -5, -2, -5}},
	}

	for i := range input {
		assert.Equal(t, utils.SliceToCycleListNode(output[i].Para1, 0),
			detectCycle(utils.SliceToCycleListNode(input[i].Para1, input[i].Pos)))
	}
}
