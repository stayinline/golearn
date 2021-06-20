package main

func maximalRectangle(matrix [][]byte) (ret int) {
	if len(matrix) == 0 {
		return 0
	}
	n := len(matrix)         //行数
	m := len(matrix[0])      //列数
	left := make([][]int, n) //记录matrix[i][j]左边的连续1的个数

	for i := 0; i < n; i++ {
		left[i] = make([]int, m) // 初始化为0
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' { //是0就跳过
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1 //是1才加和
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			width := left[i][j]           // 以matrix[i][j]为右下角的矩形的宽度
			area := width                 // 以matrix[i][j]为右下角的矩形的面积
			for k := i - 1; k >= 0; k-- { //以matrix[i][j]为右下角,所以k=i-1,并且递减
				width = getMin(width, left[k][j])
				area = getMax(area, width*(i-k+1))
			}
			ret = getMax(ret, area)
		}
	}
	return
}

func main() {

}
