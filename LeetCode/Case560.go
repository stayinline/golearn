package main

import "fmt"

func subarraySum(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		tmpSun := 0
		for j := i; j < len(nums); j++ {
			tmpSun += nums[j]
			if tmpSun == k {
				ret++
			}
		}
	}
	return ret
}

func subarraySum2(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		tmpSun := 0
		for j := i; j >= 0; j-- {
			tmpSun += nums[j]
			if tmpSun == k {
				ret++
			}
		}
	}
	return ret
}

func main() {
	nums := []int{1, 2, 3, 1, 2}
	i := subarraySum(nums, 6)
	fmt.Println(i)

}
