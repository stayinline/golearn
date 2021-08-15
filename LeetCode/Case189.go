package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[(i+k)%n] = nums[i]
	}
	copy(nums, tmp)
}

func rotate2(nums []int, k int) {
	n := len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[n-1]
		for j := n - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = tmp
	}
}

func main() {

	nums := []int{1, 2}
	rotate2(nums, 3)

	fmt.Println(nums)

}
