package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	numsCopy := make([]int, n)
	copy(numsCopy, nums)
	sort.Ints(numsCopy)

	i, j := 0, n-1
	for i < n {
		if nums[i] == numsCopy[i] {
			i++
		} else {
			break
		}
	}
	fmt.Println("i=", i)
	for j >= 0 {
		if numsCopy[j] == nums[j] {
			j--
		} else {
			break
		}
	}
	fmt.Println("j=", j)
	if j-i >= 0 {
		return j - i + 1
	} else {
		return 0
	}
}

func findUnsortedSubarray2(nums []int) int {
	min, max, n := 0, 0, len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			min = nums[i+1]
			fmt.Println("min=", min)
			break
		}
	}

	for i := n - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			max = nums[i-1]
			fmt.Println("max=", max)
			break
		}
	}

	l, r := 0, n-1
	for ; l < n; l++ {
		if nums[l] > min {
			break
		}
	}

	for ; r > 0; r-- {
		if nums[r] < max {
			break
		}
	}
	fmt.Println("l=", l)
	fmt.Println("r=", r)

	if r-l >= 0 {
		return r - l + 1
	} else {
		return 0
	}

}

func main() {
	//nums := []int{2, 6, 4, 8, 10, 9, 15}
	nums2 := []int{1, 2, 3, 4}
	//ret := findUnsortedSubarray(nums)
	//fmt.Println(ret)

	ret2 := findUnsortedSubarray2(nums2)
	fmt.Println(ret2)

}
