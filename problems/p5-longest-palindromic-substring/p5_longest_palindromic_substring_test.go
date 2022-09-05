package longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type data struct {
		input    string
		excepted string
	}

	testData := []data{
		{
			input:    "babad",
			excepted: "bab",
		},
		{
			input:    "cbbd",
			excepted: "bb",
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, longestPalindrome(tdata.input))
	}
}
