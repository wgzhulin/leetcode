package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：n = 2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1

示例 2：
输入：n = 3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2

示例 3：
输入：n = 4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3
*/
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
