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

func IntSliceToBinaryTree(input []int) *basedata.TreeNode {
	s := make([]int, len(input)+1)
	copy(s[1:], input)
	return binaryTree(s, 1)
}

func binaryTree(input []int, index int) *basedata.TreeNode {
	if index < len(input) {
		val := input[index]
		if val != 0 {
			return &basedata.TreeNode{
				Val:   input[index],
				Left:  binaryTree(input, index*2),
				Right: binaryTree(input, index*2+1),
			}
		}
	}
	return nil
}
