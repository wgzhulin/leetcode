package longest_palindrome

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type Input struct {
		Para1 string
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{
			Para1: "abccccdd",
		},
		{
			Para1: "a",
		},
		{
			Para1: "bb",
		},
		{
			Para1: "ccc",
		},
		{
			Para1: "aabbcc",
		},
	}

	output := []Output{
		{Para1: 7},
		{Para1: 1},
		{Para1: 2},
		{Para1: 3},
		{Para1: 6},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, longestPalindrome(input[i].Para1))
	}
}
