package code

func SliceToListNode(input []int) *ListNode {
	result := &ListNode{}

	head := result
	for i := range input {
		head.Val = input[i]
		if i < len(input) -1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return result
}
