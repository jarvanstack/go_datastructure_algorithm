package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	_ = printMatrix(matrix)
}
func printMatrix(matrix [][]int) error {
	if matrix == nil {
		fmt.Printf("%s\n", "matrix == nil")
		return fmt.Errorf("error:matrix == nil")
	}
	width := len(matrix)
	if  width <= 1 {
		return fmt.Errorf("error: widhth == 0")
	}
	length := len(matrix[0])
	if length <= 1  {
		return fmt.Errorf("error:length or widhth == 0")
	}
	start := 0
	for start <= length/2 && start <= width/2 {
		printOneCircle(matrix,length,width, start)
		start++

	}
	return nil
}
//打印一个圈的矩阵.
func printOneCircle(matrix [][]int,length int,width int,start int)  {
	endX := length - start - 1
	endY := width - start - 1
	//一步
	for i := start; i <= endX; i++ {
		fmt.Printf(" %d ", matrix[start][i])
	}
	//二步
	for i := start + 1; i <= endY; i++ {
		fmt.Printf(" %d ", matrix[i][endX])
	}
	if start < endY {
		//三步
		for i := endX - 1; i >= start; i-- {
			fmt.Printf(" %d ", matrix[endY][i])
		}
	}
	//四步
	for i := endY - 1; i >= start + 1; i-- {
		fmt.Printf(" %d ", matrix[i][start])
	}
}
