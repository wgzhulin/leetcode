package problems

import "github.com/zhulinw/leetcode/basedata"

/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeNthFromEnd(head *basedata.ListNode, n int) *basedata.ListNode {
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