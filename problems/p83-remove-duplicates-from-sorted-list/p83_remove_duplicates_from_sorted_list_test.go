package remove_duplicates_from_sorted_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 []int
	}

	input := []Input{
		{
			Para1: []int{1, 1, 2},
		},
		{
			Para1: []int{1, 1, 2, 3, 3},
		},
		{
			Para1: []int{1, 1, 1},
		},
	}

	output := []Output{
		{
			Para1: []int{1, 2},
		},
		{
			Para1: []int{1, 2, 3},
		},
		{
			Para1: []int{1},
		},
	}

	for i := range input {
		assert.Equal(t, utils.SliceToListNode(output[i].Para1),
			deleteDuplicates(utils.SliceToListNode(input[i].Para1)))
	}
}
