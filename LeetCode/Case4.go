package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, num := range nums2 {
		nums1 = append(nums1, num)
	}

	sort.Ints(nums1)
	fmt.Println(nums1)
	l := len(nums1)
	if l%2 == 0 {
		tmp := float64(nums1[l/2]+nums1[l/2-1]) / 2
		return tmp
	} else {
		return float64(nums1[l/2])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	nums := make([]int, l1+l2)
	j1, k2 := 0, 0
	for i := 0; i < l1+l2; i++ {
		if j1 < l1 && k2 < l2 {
			if nums1[j1] < nums2[k2] {
				nums[i] = nums1[j1]
				j1++
			} else {
				nums[i] = nums2[k2]
				k2++
			}
		} else if j1 < l1 {
			nums[i] = nums1[j1]
			j1++
		} else {
			nums[i] = nums2[k2]
			k2++
		}
	}
	idx := (l1 + l2) / 2
	if (l1+l2)%2 == 0 {
		tmp := float64(nums[idx]+nums[idx-1]) / 2
		return tmp
	} else {
		return float64(nums[idx])
	}
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			return nums[mid]
		}
	}
	return 0
}

func main() {
	nums1 := []int{0, 0, 0, 0, 0}
	nums2 := []int{-1, 0, 0, 0, 0, 0, 1}
	findMedianSortedArrays2(nums1, nums2)

}
