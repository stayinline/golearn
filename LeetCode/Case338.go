package main

import "fmt"

func countBits(n int) (res []int) {
	res = make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = getBitsOfN(i)
	}
	return
}

func getBitsOfN(n int) (count int) {
	count = 0
	for n > 0 {
		count++
		n &= n - 1
	}
	return
}

func main() {

	fmt.Println(countBits(5))

}
