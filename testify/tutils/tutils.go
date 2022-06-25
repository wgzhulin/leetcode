package tutils

import "github.com/zhulinw/leetcode/basedata"

func SliceToListNode(input []int) *basedata.ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &basedata.ListNode{}
	head := result
	for i := range input {
		head.Val = input[i]
		if i < len(input)-1 {
			head.Next = &basedata.ListNode{}
			head = head.Next
		}
	}
	return result
}
