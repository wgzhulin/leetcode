package problems

import (
	. "github.com/zhulinw/leetcode/ltdata"
)

/*
21. 合并两个有序链表
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-two-sorted-lists/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	for node1, node2 := list1, list2; node1 != nil || node2 != nil; {
		if node1 != nil && node2 != nil {
			if node1.Val < node2.Val {
				head.Next = &ListNode{Val: node1.Val}
				node1 = node1.Next

				head = head.Next
			} else if node2.Val < node1.Val {
				head.Next = &ListNode{Val: node2.Val}
				node2 = node2.Next

				head = head.Next
			} else {
				tempNode := &ListNode{Val: node1.Val}
				tempNode.Next = &ListNode{Val: node2.Val}
				head.Next = tempNode

				head = tempNode.Next

				node1 = node1.Next
				node2 = node2.Next
			}
		}

		if node1 != nil && node2 == nil {
			head.Next = &ListNode{Val: node1.Val}
			node1 = node1.Next
			head = head.Next
		}

		if node1 == nil && node2 != nil {
			head.Next = &ListNode{Val: node2.Val}
			node2 = node2.Next
			head = head.Next
		}
	}
	return result.Next
}
