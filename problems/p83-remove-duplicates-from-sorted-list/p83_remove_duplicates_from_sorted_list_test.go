package remove_duplicates_from_sorted_list

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
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
		assert.Equal(t, tutils.SliceToListNode(output[i].Para1),
			deleteDuplicates(tutils.SliceToListNode(input[i].Para1)))
	}
}
