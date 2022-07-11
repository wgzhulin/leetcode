package problems

/*
27. 移除元素
给你一个数组nums 和一个值val，你需要原地移除所有数值等于 val 的元素，并返回移除后数组的新长度。不要使用额外的数组空间，你必须仅使用O(1)额外空间并原地修改输入数组。元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。 说明:为什么返回数值是整数，但输出的答案是数组呢?请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。你可以想象内部操作如下://nums是以“引用”方式传递的。也就是说，不对实参作任何拷贝intlen=removeElement(nums,val);//在函数里修改输入数组对于调用者是可见的。//根据你的函数返回的长度,它会打印出数组中该长度范围内的所有元素。for(inti=0;i

示例1：
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]

示例2：
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]


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
