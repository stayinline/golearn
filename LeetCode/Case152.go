package main

import "fmt"

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = getMax2(mx*nums[i], getMax2(nums[i], mn*nums[i]))
		minF = getMin2(mn*nums[i], getMin2(nums[i], mx*nums[i]))
		ans = getMax2(maxF, ans)
	}
	return ans
}

func getMax2(a1 int, a2 int) int {
	if a1 > a2 {
		return a1
	} else {
		return a2
	}
}

func getMin2(a1 int, a2 int) int {
	if a1 > a2 {
		return a2
	} else {
		return a1
	}
}

func main() {
	nums := []int{2, 3, -2, 4}
	product := maxProduct(nums)
	fmt.Println(product)

}
