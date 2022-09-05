package isomorphic_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	type Input struct {
		Para1 string
		Para2 string
	}

	type Output struct {
		Para1 bool
	}

	input := []Input{
		{
			Para1: "egg",
			Para2: "add",
		},
		{
			Para1: "foo",
			Para2: "bar",
		},
		{
			Para1: "paper",
			Para2: "title",
		},
		{
			Para1: "badc",
			Para2: "baba",
		},
	}

	output := []Output{
		{Para1: true},
		{Para1: false},
		{Para1: true},
		{Para1: false},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, isIsomorphic2(input[i].Para1, input[i].Para2))
	}
}
