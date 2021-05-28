package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	i, j := 0, 0
	for ; j < n; {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	fmt.Println(nums)
}

func main() {

	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

}
