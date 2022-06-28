package problems

import (
	"github.com/zhulinw/leetcode/basedata"
	"math"
)

/*
515. 在每个树行中找最大值
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。

示例1：
输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]

示例2：
输入: root = [1,2,3]
输出: [1,3]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func largestValues(root *basedata.TreeNode) []int {
	ans := make([]int, 0)

	queue := make([]*basedata.TreeNode, 0)
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

func largestValues2(root *basedata.TreeNode) []int {
	return dfsTreeNode(root, []int{}, 0)
}

func dfsTreeNode(root *basedata.TreeNode, ans []int, height int) []int {
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
