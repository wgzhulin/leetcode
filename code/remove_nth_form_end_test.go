package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]
*/
func TestRemoveNthFromEnd(t *testing.T) {
	type data struct {
		input  []int
		input2 int
		except []int
	}

	testData := []data{
		{
			input:  []int{1, 2, 3, 4, 5},
			input2: 2,
			except: []int{1, 2, 3, 5},
		},
		{
			input:  []int{1},
			input2: 1,
			except: []int{},
		},
		{
			input:  []int{1, 2},
			input2: 1,
			except: []int{1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, SliceToListNode(tdata.except),
			removeNthFromEnd(SliceToListNode(tdata.input), tdata.input2))
	}
}
