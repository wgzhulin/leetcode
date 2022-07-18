package ltdata

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (l *ListNode) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for node := l; node != nil; node = node.Next {
		builder.WriteString(fmt.Sprintf("%v", node.Val))
		if node.Next != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")

	return builder.String()
}
