package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
*/
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
		assert.Equal(t, tdata.except, tdata.input)
	}
}
