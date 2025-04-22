package int_algo

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 顺时针螺旋打印N列二维数组，任务

输入：matrix =
[
[1,2,3],
[4,5,6],
[7,8,9]
]
输出：[1,2,3,6,9,8,7,4,5]
*/

// 官方题解
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

/*
变形：给定任意一个位置，返回顺时针螺旋顺序的元素，
*/
func GenerateListV2(list [][]int, startRow, startCol int) []int {
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

//自己的解法
/*
[
[1,2,3],
[4,5,6],
[7,8,9]
]
1 2

*/
// 自己写的SpiralOrderV2
func SpiralOrderV2(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	leftCol, rightCol := 0, cols-1
	topRow, bottomRow := 0, rows-1
	//res := make([]int, 0, rows*cols) 这样没有给了容量，但是没有赋予0值，不会有问题；
	res := make([]int, 0) //leng

	for leftCol <= rightCol && topRow <= bottomRow {
		// 从左边到右边，遍历第一行；其实就是第一列到最后一列的第一层数字，其实是遍历列
		//从左往右遍历topRow
		for col := leftCol; col <= rightCol; col++ {
			res = append(res, matrix[topRow][col])
		}
		topRow++
		// 从上边到下边的最后一列，其实，就是从第一行到最后一行，所以是遍历行
		//从上往下遍历rightCol
		for row := topRow; row <= bottomRow; row++ {
			res = append(res, matrix[row][rightCol])
		}
		rightCol--
		// 从右到左边，需要判断是否还有行,遍历行
		if bottomRow >= topRow {
			for col := rightCol; col >= leftCol; col-- {
				res = append(res, matrix[bottomRow][col])
			}
			bottomRow--
		}

		// 从下往上，遍历列；
		//同时判断，是否存在列；
		if leftCol <= rightCol {
			for row := bottomRow; row > topRow; row-- {
				res = append(res, matrix[row][leftCol])
			}
			leftCol++
		}
	}

	return res
}

/*  1为啥最后两个循环需要加if bottomRow >= topRow ？{
  		是因为前面已经对于topRow，rightCol进行处理了；
    2为啥前面两个加？
 		是因为总的循环里面已经加了，第一个for循环一定是合理的，不用显式的再加了，
    3那为啥第二个也不需要加？，他已经对于topRow进行了变化了呀？
		是因为他的逻辑里面已经是对于row的for循环；（从上往下遍历rightCol） for row := topRow; row <= bottomRow; row++ 已经做了这件事了，没必要加了；


总结来说，
	1这里为啥后面加，前面不加，是为了简洁，前面的没必要加
换个思路：
	2后面两个加的是if，没有直接break，是因为后面的两种情况，有可能其中一个不满足，另外一个是满足的；
*/

/*
func main() {
    testCases := []struct {
        matrix [][]int
        want   []int
    }{
        {
            matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
            want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
        },
        {
            matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
            want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
        },
        {
            matrix: [][]int{{1}},
            want:   []int{1},
        },
    }

    for _, tc := range testCases {
        got := spiralOrderV2(tc.matrix)
        fmt.Printf("Input: %v\nOutput: %v\nExpected: %v\n\n", tc.matrix, got, tc.want)
    }
}

*/
