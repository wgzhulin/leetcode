package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type data struct {
		input  []string
		except string
	}

	testData := []data{
		{
			input:  []string{"flower", "flow", "flight"},
			except: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			except: "",
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, longestCommonPrefix(tdata.input))
	}
}
