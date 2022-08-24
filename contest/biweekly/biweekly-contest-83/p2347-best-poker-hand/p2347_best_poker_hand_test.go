package best_poker_hand

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestBestHand(t *testing.T) {
	type Input struct {
		Para1 []int
		Para2 []byte
	}

	type Output struct {
		Para1 string
	}

	input := []Input{
		{
			Para1: []int{13, 2, 3, 1, 9},
			Para2: []byte{'a', 'a', 'a', 'a', 'a'},
		},
		{
			Para1: []int{4, 4, 2, 4, 4},
			Para2: []byte{'d', 'a', 'a', 'b', 'c'},
		},
		{
			Para1: []int{10, 10, 2, 12, 9},
			Para2: []byte{'a', 'b', 'c', 'a', 'd'},
		},
	}

	output := []Output{
		{Para1: "Flush"},
		{Para1: "Three of a Kind"},
		{Para1: "Pair"},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, bestHand(input[i].Para1, input[i].Para2))
	}
}
