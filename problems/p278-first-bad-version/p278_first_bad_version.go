package first_bad_version

/*
278. 第一个错误的版本

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/first-bad-version/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

var stub int

func isBadVersion(n int) bool {
	if n == stub {
		return true
	}
	return false
}
