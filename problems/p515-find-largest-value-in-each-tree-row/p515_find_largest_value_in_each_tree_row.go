package find_largest_value_in_each_tree_row

import (
	. "github.com/zhulinw/leetcode/ltdata"
	"math"
)
/*
515. 在每个树行中找最大值

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		maxVal := math.MinInt32
		tmp := queue
		queue = nil

		for _, node := range tmp {
			maxVal = MaxInt(maxVal, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, maxVal)
	}
	return ans
}

func largestValues2(root *TreeNode) []int {
	return dfsTreeNode(root, []int{}, 0)
}

func dfsTreeNode(root *TreeNode, ans []int, height int) []int {
	if root == nil {
		return ans
	}

	if len(ans) == height {
		ans = append(ans, root.Val)
	} else {
		ans[height] = MaxInt(ans[height], root.Val)
	}

	if root.Left != nil {
		ans = dfsTreeNode(root.Left, ans, height+1)
	}

	if root.Right != nil {
		ans = dfsTreeNode(root.Right, ans, height+1)
	}
	return ans
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}