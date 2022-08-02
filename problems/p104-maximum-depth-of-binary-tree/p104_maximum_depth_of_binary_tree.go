package maximum_depth_of_binary_tree

import . "github.com/zhulinw/leetcode/ltdata"

/*
104. 二叉树的最大深度

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxDepth(root *TreeNode) int {
	var bfs func(*TreeNode, int)
	ans := 0
	bfs = func(root *TreeNode, high int) {
		if root == nil {
			return
		}

		if high > ans {
			ans = high
		}

		if root.Left != nil {
			bfs(root.Left, high+1)
		}
		if root.Right != nil {
			bfs(root.Right, high+1)
		}
	}

	bfs(root, 1)
	return ans
}
