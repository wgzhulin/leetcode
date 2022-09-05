package climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	type data struct {
		input  int
		except int
	}

	testData := []data{
		{
			input:  2,
			except: 2,
		},
		{
			input:  3,
			except: 3,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, climbStairs(tdata.input))
		assert.Equal(t, tdata.except, climbStairs2(tdata.input))
		assert.Equal(t, tdata.except, climbStairs3(tdata.input))
	}
}
