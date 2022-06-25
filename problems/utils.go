package problems

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
