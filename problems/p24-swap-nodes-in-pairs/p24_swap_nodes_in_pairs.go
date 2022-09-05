package swap_nodes_in_pairs

import . "github.com/zhulinw/leetcode/ltdata"
/*
24. 两两交换链表中的节点

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/swap-nodes-in-pairs/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func swapPairs(head *ListNode) *ListNode {
	ans := &ListNode{Next: head}
	for node := ans; node != nil && node.Next != nil && node.Next.Next != nil; {
		node1 := node.Next
		node2 := node1.Next

		node1.Next = node2.Next
		node.Next = node2
		node2.Next = node1

		node = node.Next.Next

	}
	return ans.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node1 := head.Next
	head.Next = swapPairs2(node1.Next)
	node1.Next = head

	return node1
}
