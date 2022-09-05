package path_sum

import . "github.com/zhulinw/leetcode/ltdata"
/*
112. 路径总和

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-sum/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}

