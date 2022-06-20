package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：asteroids = [5,10,-5]
输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。

示例 2：
输入：asteroids = [8,-8]
输出：[]
解释：8 和 -8 碰撞后，两者都发生爆炸。

示例 3：
输入：asteroids = [10,2,-5]
输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

示例 4：
输入：asteroids = [-2,-1,1,2]
输出：[-2,-1,1,2]
解释：-2 和 -1 向左移动，而 1 和 2 向右移动。 由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞
*/
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
