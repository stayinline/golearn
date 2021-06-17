package main

import "fmt"

/*
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/
var ret int

func findTargetSumWays(nums []int, target int) int {
	ret = 0
	findTargetSumWaysHelper(nums, 0, target, 0)
	return ret

}

func findTargetSumWaysHelper(nums []int, i int, target int, curSum int) {
	if i == len(nums) {
		if curSum == target {
			ret++
		}
		return
	}
	findTargetSumWaysHelper(nums, i+1, target, curSum+nums[i])
	findTargetSumWaysHelper(nums, i+1, target, curSum-nums[i])
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	targetSumWays := findTargetSumWays(nums, 3)
	fmt.Println(targetSumWays)

}
