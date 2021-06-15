package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	lens := 1
	step := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			//相等就继续
			continue
		} else if (nums[i] - nums[i-1]) == 1 {
			//连续就累加
			step += 1
			lens = getMax22(lens, step)
		} else {
			//不连续则更新step
			step = 1
		}
	}
	return lens
}
func getMax22(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	//1, 2, 3, 4, 100, 200
	consecutive := longestConsecutive(nums)
	fmt.Println(consecutive)

}
