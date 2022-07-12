package problems

/*
1252. 奇数值单元格的数目
给你一个mxn的矩阵，最开始的时候，每个单元格中的值都是0。另有一个二维索引数组 indices，indices[i]=[ri,ci]指向矩阵中的某个位置，其中ri和ci分别表示指定的行和列（从0开始编号）。对indices[i]所指向的每个位置，应同时执行下述增量操作：	ri行上的所有单元格，加1。	ci列上的所有单元格，加1。给你m、n和indices。请你在执行完所有 indices 指定的增量操作后，返回矩阵中奇数值单元格的数目。 

示例1：
输入：m = 2, n = 3, indices = [[0,1],[1,1]]
输出：6

示例2：
输入：m = 2, n = 2, indices = [[1,1],[0,0]]
输出：0


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
