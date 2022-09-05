package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type data struct {
		input  []int
		input2 int
		except []int
	}

	testData := []data{
		{
			input:  []int{2, 7, 11, 15},
			input2: 9,
			except: []int{0, 1},
		},
		{
			input:  []int{3, 2, 4},
			input2: 6,
			except: []int{1, 2},
		},
		{
			input:  []int{3, 3},
			input2: 6,
			except: []int{0, 1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, twoSum(tdata.input, tdata.input2))
	}
}
