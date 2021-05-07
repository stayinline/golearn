package main

import "fmt"

func hammingDistance(x int, y int) int {
	var z = x ^ y
	var ret = 0
	for z != 0 {
		ret++
		z = z & (z - 1)
	}
	return ret
}
func main() {
	distance := hammingDistance(1, 4)

	fmt.Println(distance)

}
