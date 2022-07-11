package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type data struct {
		input1   []int
		input2   int
		excepted []int
	}

	testData := []data{
		{
			input1:   []int{3, 2, 2, 3},
			input2:   3,
			excepted: []int{2, 2},
		},
		{
			input1:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			input2:   2,
			excepted: []int{0, 1, 4, 0, 3},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, len(tdata.excepted), removeElement(tdata.input1, tdata.input2))
		// assert.Equal(t, len(tdata.excepted), removeElement2(tdata.input1, tdata.input2))
	}
}
