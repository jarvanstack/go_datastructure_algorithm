package main

import "fmt"

//归并排序 go 实现
//先还分后合并.
func main() {
	array, err := start([]int{3, 4, 1, 8, 9, 2, 5, 6, 7,11})
	if err != nil {
		fmt.Printf("err=%#v\n", err)
		return
	}
	fmt.Printf("array=%#v\n", array)
}
