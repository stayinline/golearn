package main

func sortColors(nums []int) {
	p0, p1 := 0, 0 //p0用来交换0，p1用来交换1，所以p0记录的是0后面的那个元素，p1记录的是1后面那个元素
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i] //交换0
			if p0 < p1 {                          //处理最后的数组，如果剩了一个1，应该再交换一次1
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i] //交换1
			p1++
		}
	}
}

func main() {

}
