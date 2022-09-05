package find_the_index_of_the_first_occurrence_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	type data struct {
		input1    string
		input2    string
		excepted int
	}

	testData := []data{
		{
			input1:   "hello",
			input2:   "ll",
			excepted: 2,
		},
		{
			input1:   "aaaaa",
			input2:   "baa",
			excepted: -1,
		},
		{
			input1:    "a",
			input2:    "a",
			excepted: 0,
		},
		{
			input1:    "abc",
			input2:    "c",
			excepted: 2,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, strStr(tdata.input1, tdata.input2))
		assert.Equal(t, tdata.excepted, strStr2(tdata.input1, tdata.input2))
	}
}
