package main

func numIslands(grid [][]byte) int {
	lr := len(grid)
	if lr < 0 {
		return 0
	}
	lc := len(grid[0])

	ret := 0
	for i := 0; i < lr; i++ {
		for j := 0; j < lc; j++ {
			if grid[i][j] == '1' {
				ret++
				//每一次深度优先搜索结束就是一座岛屿
				dfs(grid, i, j)
			}
		}
	}

	return ret
}

//每个grid[i][j]，上下左右进行深度优先搜索，遇到1就进入，走到最后一定是0
func dfs(grid [][]byte, i int, j int) {
	lr := len(grid)
	lc := len(grid[0])
	grid[i][j] = '0'
	//下
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	//上
	if i+1 < lr && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	//左
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	//右
	if j+1 < lc && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

func main() {

}
