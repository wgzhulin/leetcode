package add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type data struct {
		input1 []int
		input2 []int
		except []int
	}

	testData := []data{
		{
			input1: []int{2, 4, 3},
			input2: []int{5, 6, 4},
			except: []int{7, 0, 8},
		},
		{
			input1: []int{0},
			input2: []int{0},
			except: []int{0},
		},
		{
			input1: []int{9, 9, 9, 9, 9, 9, 9},
			input2: []int{9, 9, 9, 9},
			except: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tutils.SliceToListNode(tdata.except), addTwoNumbers(tutils.SliceToListNode(tdata.input1), tutils.SliceToListNode(tdata.input2)))
	}
}
