package main

type queue []int

func (q *queue) Enqueue(a int) {
	*q = append(*q, a)
}

// 注意别犯浑，先用res缓存答案
func (q *queue) Dequeue() int {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Size() int {
	return len(*q)
}

// 无向图的最短路径：BFS
func minPathValue(mat [][]int) int {
	rowQ := queue{}
	colQ := queue{}

	// visited (不仅作为marked数组，还记录路径长度)
	visited := make([][]int, len(mat))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(mat[0]))
	}

	// enqueue source node
	rowQ.Enqueue(0)
	colQ.Enqueue(0)
	visited[0][0] = 1

	var row, col int

	walkTo := func(pre, newR, newC int) {
		// array out of bounds
		if newR < 0 || newR >= len(mat) || newC < 0 || newC >= len(mat[0]) {
			return
		}
		// if visited
		if visited[newR][newC] != 0 {
			return
		}
		// if not 1
		if mat[newR][newC] != 1 {
			return
		}

		visited[newR][newC] = pre + 1
		rowQ.Enqueue(newR)
		colQ.Enqueue(newC)
	}

	for !rowQ.IsEmpty() {
		row, col = rowQ.Dequeue(), colQ.Dequeue()
		// find answer, early stop
		if row == len(mat)-1 && col == len(mat[0])-1 {
			//fmt.Println(visited)
			return visited[len(mat)-1][len(mat[0])-1]
		}
		// find (row, col) neighbours
		walkTo(visited[row][col], row-1, col)
		walkTo(visited[row][col], row+1, col)
		walkTo(visited[row][col], row, col-1)
		walkTo(visited[row][col], row, col+1)
	}
	return -1 // cannot reach
}

// 当前节点不判断visited，而是判断其子节点，如果visited就不会进入queue
// 类似层序遍历的方式，记录层数
func minPathValue2(mat [][]int) int {
	rowQ := queue{}
	colQ := queue{}

	// visited (不仅作为marked数组，还记录路径长度)
	visited := make([][]int, len(mat))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(mat[0]))
	}

	// enqueue source node
	rowQ.Enqueue(0)
	colQ.Enqueue(0)
	visited[0][0] = 1

	walkTo := func(path, newR, newC int) {
		// array out of bounds
		if newR < 0 || newR >= len(mat) || newC < 0 || newC >= len(mat[0]) {
			return
		}
		// if visited
		if visited[newR][newC] != 0 {
			return
		}
		// if not 1
		if mat[newR][newC] != 1 {
			return
		}

		visited[newR][newC] = 1
		rowQ.Enqueue(newR)
		colQ.Enqueue(newC)
	}

	var row, col int
	var path int
	for !rowQ.IsEmpty() {
		size := rowQ.Size()
		path += 1
		for i := 0; i < size; i++ {
			row, col = rowQ.Dequeue(), colQ.Dequeue()

			if row == len(mat)-1 && col == len(mat[0])-1 {
				return path
			}
			// find (row, col) neighbours
			walkTo(path, row-1, col)
			walkTo(path, row+1, col)
			walkTo(path, row, col-1)
			walkTo(path, row, col+1)
		}
	}
	return -1
}

//func main() {
//	mat := [][]int{
//		{1, 1, 1, 1, 1},
//		{1, 0, 1, 0, 1},
//		{1, 1, 1, 0, 1},
//		{0, 0, 0, 0, 1},
//	}
//
//	fmt.Println(minPathValue(mat))
//	fmt.Println(minPathValue2(mat))
//}
