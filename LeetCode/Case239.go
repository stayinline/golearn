package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) (ret []int) {
	queue := make([]int, k)

	for j, i := 0, 1-k; j < len(nums); {
		if i > 0 && queue[0] == nums[i-1] {
			queue = queue[1:]
		}

		for len(queue) != 0 && queue[len(queue)-1] < nums[j] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[j])
		fmt.Println(queue)
		if i >= 0 {
			ret = append(ret, queue[0])
		}
		i++
		j++
	}
	return
}

// 超时的暴力法
func maxSlidingWindow1(nums []int, k int) (ret []int) {
	if k > len(nums) {
		return nums
	}
	for i, j := 0, k-1; j < len(nums); {
		max := nums[i]
		for _, a := range nums[i : j+1] {
			max = getMax21(a, max)
		}
		ret = append(ret, max)
		i++
		j++
	}
	return
}
func getMax21(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxSlidingWindowOfficial(nums []int, k int) []int {
	var q []int
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	ret := maxSlidingWindow(nums, 3)
	fmt.Println(ret)
	//fmt.Println(nums[0:2])
	//fmt.Println(nums[:len(nums)-1])

}
