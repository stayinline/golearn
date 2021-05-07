package main

import "fmt"

//判断 s 是否为 t 的子序列。
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	sl, tl := len(s), len(t)
	for i < sl && j < tl {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == sl
}

func main() {
	s := "abc"
	t := "aqqqbc"
	ret := isSubsequence(s, t)
	println(ret)
	fmt.Println(ret)

}
