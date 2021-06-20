package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	fmt.Println(dp)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = getMin21(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	fmt.Println(dp)
	return dp[n]
}

func getMin21(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {

	numSquares(12)

}
