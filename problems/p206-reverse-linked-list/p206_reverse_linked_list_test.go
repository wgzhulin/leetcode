package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
	"testing"
)

func TestReverseList(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 []int
	}

	input := []Input{
		{
			Para1: []int{1, 2, 3, 4, 5},
		},
		{
			Para1: []int{1, 2},
		},
	}

	output := []Output{
		{Para1: []int{5, 4, 3, 2, 1}},
		{Para1: []int{2, 1}},
	}

	for i := range input {
		assert.Equal(t, tutils.SliceToListNode(output[i].Para1), reverseList2(tutils.SliceToListNode(input[i].Para1)))
	}
}
