package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	type data struct {
		input  string
		except int
	}

	testData := []data{
		{
			input:  "Hello World",
			except: 5,
		},
		{
			input:  "   fly me   to   the moon  ",
			except: 4,
		},
		{
			input:   "luffy is still joyboy",
			except: 6,
		},
		{
			input:   "a",
			except: 1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, lengthOfLastWord2(tdata.input))
	}
}
