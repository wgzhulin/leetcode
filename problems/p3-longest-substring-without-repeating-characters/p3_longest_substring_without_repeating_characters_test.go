package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type data struct {
		input  string
		except int
	}

	testData := []data{
		{
			input:  "abcabcbb",
			except: 3,
		},
		{
			input:  "bbbbb",
			except: 1,
		},
		{
			input:  "pwwkew",
			except: 3,
		},
		{
			input:  " ",
			except: 1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, lengthOfLongestSubstring(tdata.input))
	}
}
