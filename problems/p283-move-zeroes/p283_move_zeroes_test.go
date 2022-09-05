package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type data struct {
		input  []int
		except []int
	}

	testData := []data{
		{
			input:  []int{0, 1, 0, 3, 12},
			except: []int{1, 3, 12, 0, 0},
		},
		{
			input:  []int{0},
			except: []int{0},
		},
	}

	for _, tdata := range testData {
		moveZeroes(tdata.input)
		moveZeroes2(tdata.input)
		assert.Equal(t, tdata.except, tdata.input)
	}
}
