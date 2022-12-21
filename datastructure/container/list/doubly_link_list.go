package list

import . "github.com/zhulinw/leetcode/ltdata"

type doublyLinkList struct {
	size int
	Head *DoublyNode
}

func NewDoublyLinkList() *doublyLinkList {
	return &doublyLinkList{Head: &DoublyNode{}}
}

func (l *doublyLinkList) Insert(n int, v int) {
}

func (l *doublyLinkList) Remove(n int) int {
}
