package remove_nth_node_from_end_of_list

import . "github.com/zhulinw/leetcode/ltdata"
/*
19. 删除链表的倒数第 N 个结点

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sum := 0
	for node := head; node != nil; node = node.Next {
		sum++
	}

	if sum-n-1 < 0 {
		head = head.Next
		return head
	}

	cur := 0
	for node := head; node != nil; node = node.Next {
		if cur == sum-n-1 {
			node.Next = node.Next.Next
			break
		}
		cur++
	}

	return head
}
