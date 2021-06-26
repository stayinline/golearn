package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(heights)
	fmt.Println(area)

}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights)
	//arr := make([]int, n)
	for i := 0; i < n; i++ {
		curHeight, width := heights[i], 1
		for left := i - 1; left >= 0 && heights[left] >= curHeight; left-- {
			width++
		}
		for right := i + 1; right < n && heights[right] >= curHeight; right++ {
			width++
		}
		//fmt.Println(width)
		maxArea = getMax23(maxArea, width*curHeight)
		//arr[i] = maxArea
	}
	//fmt.Println(arr)
	return maxArea
}
func getMax23(x, y int) int {
	if x > y {
		return x
	}
	return y
}
