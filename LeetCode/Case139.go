package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, dict := range wordDict {
		wordDictSet[dict] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for end := 1; end <= len(s); end++ {
		for start := 0; start < end; start++ {
			fmt.Println(s[start:end])
			if dp[start] && wordDictSet[s[start:end]] {
				dp[end] = true
				fmt.Println(true)
				break
			}
		}
		fmt.Println()
	}
	return dp[len(s)]
}

func main() {
	s := "leetcodelet"
	words := []string{"leet", "let", "code"}
	wordBreak(s, words)

}
