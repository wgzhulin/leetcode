package asteroid_collision

/*
735. 行星碰撞

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/asteroid-collision/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return nil
	}

	for i := 1; i < len(asteroids); i++ {
		if asteroids[i-1] >= 0 && asteroids[i] < 0 {
			// 碰撞，向左移动行星质量大，会继续碰撞之前的行星
			collision(asteroids, i-1, i)
		}
	}
	return FilterZero(asteroids)
}

// 两行星碰撞后，质量较小的行星爆炸（置为0）
func collision(asteroids []int, i, j int) {
	for index := i; index >= 0; index-- {
		if asteroids[index] == 0 {
			continue
		}
		left, right := asteroids[index], asteroids[j]
		if left*right < 0 && left+right > 0 {
			// 向右移动的行星比向左移动的行星的大
			asteroids[j] = 0
			break
		} else if left*right < 0 && left+right < 0 {
			// 继续向左碰撞
			asteroids[index] = 0
		} else if left+right == 0 {
			asteroids[index], asteroids[j] = 0, 0
			break
		}
	}
}

func FilterZero(input []int) []int {
	result := make([]int, 0, len(input))
	for i := range input {
		if input[i] != 0 {
			result = append(result, input[i])
		}
	}
	return result
}