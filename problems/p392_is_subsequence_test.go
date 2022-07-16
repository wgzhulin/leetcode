package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	type Input struct {
		Para1 string
		Para2 string
	}

	type Output struct {
		Para1 bool
	}

	input := []Input{
		{
			Para1: "abc",
			Para2: "ahbgdc",
		},
		{
			Para1: "axc",
			Para2: "ahbgdc",
		},
	}

	output := []Output{
		{Para1: true},
		{Para1: false},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, isSubsequence2(input[i].Para1, input[i].Para2))
	}
}
