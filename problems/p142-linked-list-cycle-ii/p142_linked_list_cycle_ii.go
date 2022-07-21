package linked_list_cycle_ii

import (
	. "github.com/zhulinw/leetcode/ltdata"
)

/*
142. 环形链表 II

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/linked-list-cycle-ii/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func detectCycle(head *ListNode) *ListNode {
	s := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		s = append(s, node)

		for i := range s {
			if s[i] == node.Next {
				return s[i]
			}
		}
	}
	return nil
}
