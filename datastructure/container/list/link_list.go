package list

import . "github.com/zhulinw/leetcode/ltdata"

type linkList struct {
	size int
	Head *ListNode
}

func NewLinkList() *linkList {
	return &linkList{Head: &ListNode{}}
}

func (l *linkList) Insert(n int, v int) {
	index := 0
	for node := l.Head; node != nil; node = node.Next {
		if index == n {
			newNode := &ListNode{
				Val:  v,
				Next: node.Next,
			}
			node.Next = newNode
			l.size++
			break
		}
		index++
	}
}

func (l *linkList) PushFront(v int) {
	l.Insert(0, v)
}

func (l *linkList) PushBack(v int) {
	l.Insert(l.size, v)
}

func (l *linkList) PopFront() int {
	return l.Remove(0)
}

func (l *linkList) PopBack() int {
	if l.size == 0 {
		panic("no element")
	}
	return l.Remove(l.size - 1)
}

func (l *linkList) Remove(n int) int {
	if l.size == 0 {
		panic("no element")
	}
	if n > l.size || n < 0 {
		panic("out index")
	}
	index := 0
	result := -1
	for node := l.Head; node != nil; node = node.Next {
		if index == n {
			result = node.Next.Val
			node.Next = node.Next.Next

			l.size--
			break
		}
		index++
	}
	return result
}

func (l *linkList) RemoveValue(v int) {
	if l.size == 0 {
		panic("no element")
	}

	for node := l.Head; node != nil; node = node.Next {
		if node.Next != nil && node.Next.Val == v {
			node.Next = node.Next.Next
		}
	}
}

func (l *linkList) Size() int {
	return l.size
}

func (l *linkList) IsEmpty() bool {
	return l.size == 0
}

func (l *linkList) ValueAtIndex(n int) int {
	index := 0
	for node := l.Head.Next; node != nil; node = node.Next {
		if index == n {
			return node.Val
		}
		index++
	}
	panic("not found")
}

func (l *linkList) Reverse() *linkList {
	head := l.Head.Next
	l.Head.Next = nil
	for node := head; node != nil; node = node.Next {
		l.PushFront(node.Val)
	}
	return l
}

func (l *linkList) show() []int {
	result := make([]int, 0, l.size)
	for node := l.Head.Next; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return result
}
