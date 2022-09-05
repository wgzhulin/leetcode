package cells_with_odd_values_in_a_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddCells(t *testing.T) {
	type data struct {
		input1 int
		input2 int
		input3 [][]int
		except int
	}

	testData := []data{
		{
			input1: 2,
			input2: 3,
			input3: [][]int{{0, 1}, {1, 1}},
			except: 6,
		},
		{
			input1: 2,
			input2: 2,
			input3: [][]int{{1, 1}, {0, 0}},
			except: 0,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, oddCells(tdata.input1, tdata.input2, tdata.input3))
	}
}
