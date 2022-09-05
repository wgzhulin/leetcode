package add_two_numbers

import . "github.com/zhulinw/leetcode/ltdata"
/*
2. 两数相加

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/add-two-numbers/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
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
		head.Next = &ListNode{Val: (tempVal + carry) % 10}
		head = head.Next

		carry = (tempVal + carry) / 10
	}

	if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}

	return result.Next
}