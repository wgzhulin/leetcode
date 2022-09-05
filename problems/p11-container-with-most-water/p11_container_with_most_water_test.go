package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	type data struct {
		input  []int
		except int
	}

	testData := []data{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			except: 49,
		},
		{
			input:  []int{1, 1},
			except: 1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, maxArea(tdata.input))
	}
}
