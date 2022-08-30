package letter_combinations_of_a_phone_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type Input struct {
		Para1 string
	}

	type Output struct {
		Para1 []string
	}

	input := []Input{
		{
			Para1: "23",
		},
		{
			Para1: "",
		},
		{
			Para1: "2",
		},
	}

	output := []Output{
		{
			Para1: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			Para1: []string{},
		},
		{
			Para1: []string{"a", "b", "c"},
		},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, letterCombinations(input[i].Para1))
	}
}
