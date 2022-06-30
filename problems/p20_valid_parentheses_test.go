package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	type data struct {
		input    string
		excepted bool
	}

	testData := []data{
		{
			input:    "()",
			excepted: true,
		},
		{
			input:    "()[]{}",
			excepted: true,
		},
		{
			input:    "(]",
			excepted: false,
		},
		{
			input:    "([)]",
			excepted: false,
		},
		{
			input:    "{[]}",
			excepted: true,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, isValid(tdata.input))
	}
}
