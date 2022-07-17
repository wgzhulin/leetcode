package problems

import "github.com/zhulinw/leetcode/ltdata"

/*
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

https://leetcode.cn/problems/add-two-numbers/
*/
func addTwoNumbers(l1 *ltdata.ListNode, l2 *ltdata.ListNode) *ltdata.ListNode {
	result := &ltdata.ListNode{}
	carry := 0
	head := result

	for node1, node2 := l1, l2; node1 != nil || node2 != nil; {
		tempVal := 0
		if node1 != nil {
			tempVal += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			tempVal += node2.Val
			node2 = node2.Next
		}
		head.Next = &ltdata.ListNode{Val: (tempVal + carry) % 10}
		head = head.Next

		carry = (tempVal + carry) / 10
	}

	if carry > 0 {
		head.Next = &ltdata.ListNode{Val: carry}
	}

	return result.Next
}
