package n_ary_tree_preorder_traversal

import "github.com/zhulinw/leetcode/ltdata"

/*
589. N 叉树的前序遍历

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-ary-tree-preorder-traversal/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func preorder(root *ltdata.Node) []int {
	ans := make([]int, 0)
	return preorderRcur(root, ans)
}

func preorderRcur(root *ltdata.Node, ans []int) []int {
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)
	for _, child := range root.Children {
		ans = preorderRcur(child, ans)
	}
	return ans
}
