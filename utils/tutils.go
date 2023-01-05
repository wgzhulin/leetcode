package utils

import "github.com/zhulinw/leetcode/ltdata"

func SliceToListNode(input []int) *ltdata.ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &ltdata.ListNode{}
	head := result
	for i := range input {
		head.Val = input[i]
		if i < len(input)-1 {
			head.Next = &ltdata.ListNode{}
			head = head.Next
		}
	}
	return result
}

func SliceToCycleListNode(input []int, pos int) *ltdata.ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &ltdata.ListNode{}
	head := result

	var cycle *ltdata.ListNode
	var tail *ltdata.ListNode
	for i := range input {
		node := ltdata.NewListNode(input[i])
		if i == pos {
			cycle = node
		}
		if i == len(input)-1 {
			tail = node
		}
		head.Next = node
		head = head.Next
	}
	if tail != nil {
		tail.Next = cycle
	}

	return result.Next
}

func IntSliceToBinaryTree(input []int) *ltdata.TreeNode {
	s := make([]int, len(input)+1)
	copy(s[1:], input)
	return binaryTree(s, 1)
}

func binaryTree(input []int, index int) *ltdata.TreeNode {
	if index < len(input) {
		val := input[index]
		if val != 0 {
			return &ltdata.TreeNode{
				Val:   input[index],
				Left:  binaryTree(input, index*2),
				Right: binaryTree(input, index*2+1),
			}
		}
	}
	return nil
}
