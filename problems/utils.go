package problems

func SliceToListNode(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &ListNode{}
	head := result
	for i := range input {
		head.Val = input[i]
		if i < len(input)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return result
}

func FilterZero(input []int) []int {
	result := make([]int, 0, len(input))
	for i := range input {
		if input[i] != 0 {
			result = append(result, input[i])
		}
	}
	return result
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
