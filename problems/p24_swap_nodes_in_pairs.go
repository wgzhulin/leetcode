package problems

import "github.com/zhulinw/leetcode/ltdata"

/*
24. 两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/swap-nodes-in-pairs/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func swapPairs(head *ltdata.ListNode) *ltdata.ListNode {
	ans := &ltdata.ListNode{Next: head}
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

func swapPairs2(head *ltdata.ListNode) *ltdata.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node1 := head.Next
	head.Next = swapPairs2(node1.Next)
	node1.Next = head

	return node1
}
