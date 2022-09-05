package running_sum_of_1d_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunningSum(t *testing.T) {
	type data struct {
		input  []int
		except []int
	}

	testData := []data{
		{
			input:  []int{1, 2, 3, 4},
			except: []int{1, 3, 6, 10},
		},
		{
			input:  []int{1, 1, 1, 1, 1},
			except: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{1},
			except: []int{1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, runningSum(tdata.input))
	}
}
