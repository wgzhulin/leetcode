package asteroid_collision

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	type data struct {
		input1 []int
		except []int
	}

	testData := []data{
		{
			input1: []int{5, 10, -5},
			except: []int{5, 10},
		},
		{
			input1: []int{8, -8},
			except: []int{},
		},
		{
			input1: []int{10, 2, -5},
			except: []int{10},
		},
		{
			input1: []int{-2, -1, 1, 2},
			except: []int{-2, -1, 1, 2},
		},
		{
			input1: []int{-2, -2, 1, -2},
			except: []int{-2, -2, -2},
		},
		{
			input1: []int{-2, 2, -1, -2},
			except: []int{-2},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, asteroidCollision(tdata.input1))
	}
}
