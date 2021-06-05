package main

import "fmt"

func maxCoins(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	//创建一个n+2的数组，开头和末尾都填1
	arr := make([] int, n+2)
	for i := 1; i <= n; i++ {
		arr[i] = nums[i-1]
	}
	arr[0] = 1
	arr[n+1] = 1
	dp := make([][]int, n + 2)
	for i := 0; i < n + 2; i++ {
		dp[i] = make([]int, n + 2)
	}

	//从下往上遍历，i控制左边界，j控制右边界
	i, j, k := 0, 0, 0
	for i = n - 1; i >= 0; i-- {
		for j = i + 2; j <= n+1; j++ {
			//k在(i,j)范围内遍历气球，计算每选择一个气球的积分
			//一轮遍历完后，就能确定(i,j)的最大积分
			for k = i + 1; k < j; k++ {
				total := arr[i] * arr[k] * arr[j]
				total += dp[i][k] + dp[k][j]
				dp[i][j] = getMax1(dp[i][j], total)
			}
		}
	}
	return dp[0][n+1]
}

func getMax1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{3, 1, 5, 8}
	coins := maxCoins(nums)
	fmt.Println(coins)

}
