package cells_with_odd_values_in_a_matrix

/*
1252. 奇数值单元格的数目

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func oddCells(m int, n int, indices [][]int) int {
	// force
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	addCellNum(matrix, indices)

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] % 2 == 1 {
				ans++
			}
		}
	}
	return ans
}

func addCellNum(matrix, indices [][]int) {
	for i := range indices {
		row, col := indices[i][0], indices[i][1]
		for j := 0; j < len(matrix); j++ {
			for k := 0; k < len(matrix[j]); k++ {
				if row == j {
					matrix[j][k]++
				}
				if col == k {
					matrix[j][k]++
				}
			}
		}
	}
}
