package remove_nth_node_from_end_of_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

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
		assert.Equal(t, utils.SliceToListNode(tdata.except),
			removeNthFromEnd(utils.SliceToListNode(tdata.input), tdata.input2))
	}
}
