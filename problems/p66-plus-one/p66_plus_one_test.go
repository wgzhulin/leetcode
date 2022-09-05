package plus_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type data struct {
		input  []int
		except []int
	}

	testData := []data{
		{
			input:  []int{1, 2, 3},
			except: []int{1, 2, 4},
		},
		{
			input:  []int{4, 3, 2, 1},
			except: []int{4, 3, 2, 2},
		},
		{
			input:  []int{0},
			except: []int{1},
		},
		{
			input:  []int{9},
			except: []int{1, 0},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, plusOne(tdata.input))
		//assert.Equal(t, tdata.except, plusOne2(tdata.input))
	}
}
