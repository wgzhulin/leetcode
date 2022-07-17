package query_kth_smallest_trimmed_number

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestSmallestTrimmedNumbers(t *testing.T) {
	type Input struct {
		Para1 []string
		Para2 [][]int
	}

	type Output struct {
		Para1 []int
	}

	input := []Input{
		//{
		//	Para1: []string{"102", "473", "251", "814"},
		//	Para2: [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}},
		//},
		//{
		//	Para1: []string{"24", "37", "96", "04"},
		//	Para2: [][]int{{2, 1}, {2, 2}},
		//},
		{
			Para1: []string{"325240361872", "440618160237", "785744447413", "820980201297", "470082520306", "874146483840", "425300857082", "088636787077", "813218016629", "459000328006", "188683382600"},
			Para2: [][]int{{6, 7}, {4, 4}, {1, 8}, {11, 10}, {4, 8}, {11, 6}, {1, 1}, {3, 1}, {11, 10}},
		},
	}

	output := []Output{
		//{Para1: []int{2, 2, 1, 0}},
		//{Para1: []int{3, 0}},
		{Para1: []int{10,0,9,9,1,6,5,0,9}},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, smallestTrimmedNumbers(
			input[i].Para1, input[i].Para2))
	}
}
