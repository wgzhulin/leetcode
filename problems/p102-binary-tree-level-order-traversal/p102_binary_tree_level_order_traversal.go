package binary_tree_level_order_traversal

import "github.com/zhulinw/leetcode/ltdata"

/*
102. 二叉树的层序遍历

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-level-order-traversal/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func levelOrder(root *ltdata.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	levelNode := make([]*ltdata.TreeNode, 0)
	levelNode = append(levelNode, root)
	for len(levelNode) != 0 {
		levelVal := make([]int, 0, len(levelNode))

		nextLevelNode := make([]*ltdata.TreeNode, 0, len(levelNode))
		for i := range levelNode {
			levelVal = append(levelVal, levelNode[i].Val)
			if levelNode[i].Left != nil {
				nextLevelNode = append(nextLevelNode, levelNode[i].Left)
			}
			if levelNode[i].Right != nil {
				nextLevelNode = append(nextLevelNode, levelNode[i].Right)
			}
		}
		ans = append(ans, levelVal)
		levelNode = nextLevelNode
	}

	return ans
}
