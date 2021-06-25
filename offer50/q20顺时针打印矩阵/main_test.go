package main

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	_ = printMatrix(matrix)
}
func TestRobust(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		//{13, 14, 15, 16},
	}
	_ = printMatrix(matrix)
	fmt.Printf("%s\n", "")
	_ = printMatrix(nil)
	fmt.Printf("%s\n", "")
	matrix2 := [][]int{}
	_ = printMatrix(matrix2)
}