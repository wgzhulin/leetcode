package jump_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanJump(t *testing.T) {
	type data struct {
		input  []int
		except bool
	}

	testData := []data{
		{
			input:  []int{2, 3, 1, 1, 4},
			except: true,
		},
		{
			input:  []int{3, 2, 1, 0, 4},
			except: false,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, canJump(tdata.input))
		assert.Equal(t, tdata.except, canJump2(tdata.input))
	}
}
