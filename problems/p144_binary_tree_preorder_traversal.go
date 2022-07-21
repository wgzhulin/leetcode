package problems

import . "github.com/zhulinw/leetcode/ltdata"

/*
144. 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

示例 1：
输入：root = [1,null,2,3]
输出：[1,2,3]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-preorder-traversal/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var tmp []int
	tmp = append(tmp, root.Val)
	if root.Left != nil {
		tmp = append(tmp, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		tmp = append(tmp, preorderTraversal(root.Right)...)
	}
	return tmp
}
