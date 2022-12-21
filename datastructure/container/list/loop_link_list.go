package list

import . "github.com/zhulinw/leetcode/ltdata"

func NewLoopList() *linkList {
	head := &ListNode{}
	head.Next = head
	return &linkList{Head: head}
}

