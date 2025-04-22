package int_algo

// 顺时针螺旋打印N列二维数组，给定任意起点

func GenerateList(list [][]int, startRow, startCol int) []int {
	res := make([]int, 0)

	// 向右边
	// 标记已经访问过的
	// 尝试从当前方向顺序，如果没有，就换方向
	// 换方向，从右边，下，左，上，右边
	rows, cols := len(list), len(list[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndx := 0
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	x, y := startRow, startCol
	for {
		moved := false
		for i := 0; i < 4; i++ {
			dx, dy := directions[dirIndx][0], directions[dirIndx][1]
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] {
				res = append(res, list[nx][ny])
				visited[nx][ny] = true
				x, y = nx, ny
				moved = false
				break
			}
			dirIndx = (dirIndx + 1) % 4 // 切换方向
		}
		if !moved {
			break
		}
	}

	return res
}
