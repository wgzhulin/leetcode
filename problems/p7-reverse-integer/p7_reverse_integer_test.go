package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	type data struct {
		input    int
		excepted int
	}

	testData := []data{
		{
			input:    123,
			excepted: 321,
		},
		{
			input:    -123,
			excepted: -321,
		},
		{
			input:    120,
			excepted: 21,
		},
		{
			input:    0,
			excepted: 0,
		},
		{
			input:    1534236469,
			excepted: 0,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, reverse(tdata.input))
	}
}
