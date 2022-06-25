package example

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

func TestHasRepeatedCharactersInString(t *testing.T) {
	testData := []struct {
		input  string
		except bool
	}{
		{
			input:  "hello",
			except: true,
		},
		{
			input:  "ace",
			except: false,
		},
	}

	for _, data := range testData {
		assert.Equal(t, data.except, HasRepeatedCharactersInString(data.input))
		assert.Equal(t, data.except, HasRepeatedCharactersInString2(data.input))
		assert.Equal(t, data.except, HasRepeatedCharactersInString3(data.input))
	}
}
