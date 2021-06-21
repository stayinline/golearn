package main

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	maxSide := 0
	for i := 0; i < n; i++ { //初始化dp数组
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < n; i++ { //从第一个开始遍历
		for j := 1; j < m; j++ {
			if dp[i][j] == 1 { //遇到1才更新，0则跳过
				dp[i][j] = getMin(getMin(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				maxSide = getMax(maxSide, dp[i][j])
			}
		}
	}
	return maxSide * maxSide
}

func main() {

}
