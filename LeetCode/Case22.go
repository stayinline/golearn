package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) (ans []string) {
	cur := ""
	return backTrack(cur, 0, 0, n)
}

func backTrack(cur string, left int, right int, max int) (ans []string) {
	if len(cur) == max*2 {
		ans = append(ans, cur)
		fmt.Println("ans=", ans)
		return ans
	}

	if left < max {
		cur += "("
		ans = append(ans, backTrack(cur, left+1, right, max)...)
		cur = strings.TrimSuffix(cur, "(")
	}

	if right < left {
		cur += ")"
		ans = append(ans, backTrack(cur, left, right+1, max)...)
		cur = strings.TrimSuffix(cur, ")")
	}
	return ans
}

func main() {
	ans := generateParenthesis(3)
	fmt.Println(ans)
}
