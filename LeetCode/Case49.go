package main

func groupAnagrams(strs []string) [][]string {
	// k:统计每个单词中每个字母出现的次数的数组
	// v:单词
	countMap := map[[26]int][]string{}
	for _, str := range strs {
		count := [26]int{}
		for _, c := range str {
			//统计每个字母出现的次数
			count[c-'a']++
		}
		// 更新k,v
		countMap[count] = append(countMap[count], str)
	}

	ans := make([][]string, 0, len(countMap))
	//for k, v := range countMap {
	for _, v := range countMap {
		ans = append(ans, v)
	}
	return ans
}

func main() {

}
