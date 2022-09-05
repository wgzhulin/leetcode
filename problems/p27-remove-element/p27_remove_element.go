package remove_element

/*
27. 移除元素

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-element/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeElement(nums []int, val int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}

func removeElement2(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
