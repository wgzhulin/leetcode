package remove_duplicates_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type data struct {
		input    []int
		excepted int
	}

	testData := []data{
		{
			input:    []int{1, 1, 2},
			excepted: 2,
		},
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			excepted: 5,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, removeDuplicates(tdata.input))
		assert.Equal(t, tdata.excepted, removeDuplicates2(tdata.input))
	}
}
