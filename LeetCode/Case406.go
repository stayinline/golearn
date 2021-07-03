package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		// 根据每个人的身高降序排列
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	//fmt.Println(people)
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
		//fmt.Println(ans)
	}
	return
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(people)

	fmt.Println("hello world")

}
