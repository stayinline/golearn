package main

import (
	"fmt"
)

func maxArea(height []int) int {
	y := 0
	i, j := 0, len(height)-1
	for i < j {
		tmp := getMin(height[i], height[j]) * (j - i)
		y = getMax(y, tmp)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return y
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func getMin(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	fmt.Println(area)

}
