package middle_of_the_linked_list

import (
	"github.com/zhulinw/leetcode/ltdata"
)

/*
876. 链表的中间结点

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/middle-of-the-linked-list/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func middleNode(head *ltdata.ListNode) *ltdata.ListNode {
	ans := head

	n := 0
	for node := head; node != nil; node = node.Next {
		// node步进两次，ans步进一次
		if n%2 == 1 {
			ans = ans.Next
		}
		n++
	}
	return ans
}

func middleNode2(head *ltdata.ListNode) *ltdata.ListNode {
	ans := head
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		ans = ans.Next
	}
	return ans
}
