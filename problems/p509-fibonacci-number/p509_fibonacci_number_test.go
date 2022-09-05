package fibonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T) {
	type data struct {
		input  int
		except int
	}

	testData := []data{
		{
			input:  2,
			except: 1,
		},
		{
			input:  3,
			except: 2,
		},
		{
			input:  4,
			except: 3,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, fib(tdata.input))
	}
}
