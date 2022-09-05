package sum_root_to_leaf_numbers

import . "github.com/zhulinw/leetcode/ltdata"
/*
129. 求根节点到叶节点数字之和

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sum-root-to-leaf-numbers/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sumNumbers(root *TreeNode) int {
	return deepFirstSearch(root, 0)
}

func deepFirstSearch(root *TreeNode, prev int) int {
	if root == nil {
		return 0
	}

	prev = prev*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return prev
	}
	return deepFirstSearch(root.Left, prev) + deepFirstSearch(root.Right, prev)
}