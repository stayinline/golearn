package main

import "fmt"

func majorityElement(nums []int) int {
	ret, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			ret = nums[i]
		}

		if ret == nums[i] {
			count++
		} else {
			count--
		}
	}

	return ret
}

func main() {

	nums := []int{1, 2, 3, 3, 3}
	ret := majorityElement(nums)
	fmt.Println(ret)

}
