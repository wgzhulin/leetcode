package search_insert_position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	type data struct {
		input1   []int
		input2   int
		excepted int
	}

	testData := []data{
		{
			input1:   []int{1, 3, 5, 6},
			input2:  5,
			excepted: 2,
		},
		{
			input1:   []int{1, 3, 5, 6},
			input2:  2,
			excepted: 1,
		},
		{
			input1:   []int{1, 3, 5, 6},
			input2:  7,
			excepted: 4,
		},
		{
			input1:   []int{1, 3, 5, 7},
			input2:  6,
			excepted: 3,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, searchInsert2(tdata.input1, tdata.input2))
	}
}
