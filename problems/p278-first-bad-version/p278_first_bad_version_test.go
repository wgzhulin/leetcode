package first_bad_version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBadVersion(t *testing.T) {
	type Input struct {
		Para1 int
		Para2 int
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{
			Para1: 5,
			Para2: 4,
		},
		{
			Para1: 1,
			Para2: 1,
		},
	}

	output := []Output{
		{Para1: 4},
		{Para1: 1},
	}

	for i := range input {
		stub = input[i].Para2
		assert.Equal(t, output[i].Para1, firstBadVersion(input[i].Para1))
	}
}
