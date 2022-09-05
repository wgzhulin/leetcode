package container_with_most_water

/*
11. 盛最多水的容器

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/container-with-most-water/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		area := MinInt(height[i], height[j]) * (j - i)
		max = MaxInt(max, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}