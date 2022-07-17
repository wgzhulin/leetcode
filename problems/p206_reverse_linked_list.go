package problems

import "github.com/zhulinw/leetcode/ltdata"

/*
206. 反转链表

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-linked-list/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func reverseList(head *ltdata.ListNode) *ltdata.ListNode {
	var ans *ltdata.ListNode

	for node := head; node != nil; node= node.Next {
		ans = &ltdata.ListNode{Val: node.Val, Next: ans}
	}

	return ans
}

// 递归
func reverseList2(head *ltdata.ListNode) *ltdata.ListNode {
	return reverseNode(nil, head)
}

func reverseNode(ans, head *ltdata.ListNode) *ltdata.ListNode{
	if head == nil {
		return ans
	}
	ans = &ltdata.ListNode{
		Val:  head.Val,
		Next: ans,
	}
	return reverseNode(ans, head.Next)
}
