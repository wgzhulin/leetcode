package jump_game

/*
55. 跳跃游戏

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/jump-game/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 每次更新能跳到的最远距离，
// 当前位置超过最远能跳到的距离时，返回false
func canJump(nums []int) bool {
	max := 0
	for i := range nums {
		if i > max { // 当前位置超过最远能跳到的位置
			return false
		}
		// 更新能跳到的最大的位置
		max = MaxInt(max, i+nums[i])
	}

	return true
}

// 下次能跳的最远位置，当下次能跳的最远位置小于等于当前位置
func canJump2(nums []int) bool {
	maxJumpPos := 0
	for i := 0; i < len(nums)-1; { // pos = len(nums)-1
		nextJumpPos := getMaxJumpPos(i, nums)
		if nextJumpPos <= maxJumpPos {
			return false
		}
		i = nextJumpPos
	}

	return true
}

func getMaxJumpPos(start int, nums []int) int {
	maxJumpPos := 0

	end := start + nums[start]
	if end > len(nums) {
		end = len(nums)
	}
	for i := start; i < end; i++ {
		if i+nums[i] > maxJumpPos {
			maxJumpPos = i + nums[i]
		}
	}
	return maxJumpPos
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
