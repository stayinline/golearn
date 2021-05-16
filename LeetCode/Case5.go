package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		//找到以 s[i] 为中心的回文串，因为回文串可能是偶数个
		l1, r1 := longestSubStr(s, i, i)
		//找到以 s[i] 和 s[i+1] 为中心的回文串，因为回文串可能是奇数个
		l2, r2 := longestSubStr(s, i, i+1)
		//谁的回文串更长便选择谁
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]

}

// s在[l,r]区间内的最长回文子串的结束位置
func longestSubStr(s string, l int, r int) (int, int) {
	//l >= 0 && r < len(s) 这两个是边界条件
	//s[l] == s[r]是回文串判断条件
	//所以三者要 &&
	for ; l >= 0 && r < len(s) && s[l] == s[r]; {
		l--
		r++
	}
	return l + 1, r - 1
}

func main() {
	s := "bb"
	fmt.Println(longestPalindrome(s))
	//fmt.Println(longestSubStr(s, 2, 2))
	//fmt.Println(longestSubStr(s, 2, 3))
}
