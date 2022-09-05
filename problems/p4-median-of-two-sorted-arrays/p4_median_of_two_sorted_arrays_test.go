package median_of_two_sorted_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type data struct {
		input1 []int
		input2 []int
		except float64
	}

	testData := []data{
		{
			input1: []int{1, 3},
			input2: []int{2},
			except: 2.00000,
		},
		{
			input1: []int{1, 2},
			input2: []int{3, 4},
			except: 2.50000,
		},
		{
			input1: []int{},
			input2: []int{1, 2, 3, 4, 5},
			except: 3,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, findMedianSortedArrays(tdata.input1, tdata.input2))
	}
}
