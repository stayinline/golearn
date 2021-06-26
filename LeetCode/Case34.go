package main

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	i, j := 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			a, b := mid, mid
			for a >= 0 && nums[a] == target {
				a--
			}
			for b < len(nums) && nums[b] == target {
				b++
			}
			ret = []int{a + 1, b - 1}
			return ret
		} else if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		}
	}
	return ret
}

func main() {

}
