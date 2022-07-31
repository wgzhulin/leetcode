package remove_duplicates_from_sorted_list

import . "github.com/zhulinw/leetcode/ltdata"

/*
83. 删除排序链表中的重复元素

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func deleteDuplicates(head *ListNode) *ListNode {
	for node := head; node != nil; {
		if node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
