package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	type data struct {
		input1 []int
		input2 int
		except int
	}

	testData := []data{
		{
			input1: []int{-1, 0, 3, 5, 9, 12},
			input2: 9,
			except: 4,
		},
		{
			input1: []int{-1, 0, 3, 5, 9, 12},
			input2: 2,
			except: -1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, search(tdata.input1, tdata.input2))
	}
}
