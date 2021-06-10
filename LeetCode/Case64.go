package main

import "fmt"

func minPathSum(grid [][]int) int {

	raw, col := len(grid), len(grid[0])
	dp := make([][]int, raw)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	//fmt.Println(dp)
	for i := 1; i < col; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	fmt.Println(dp)
	for i := 1; i < raw; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	fmt.Println(dp)
	for i := 1; i < raw; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = grid[i][j] + getMin22(dp[i-1][j], dp[i][j-1])
		}
	}
	fmt.Println(dp)

	return dp[raw-1][col-1]
}

func getMin22(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(len(grid))

	fmt.Println(minPathSum(grid))

}
