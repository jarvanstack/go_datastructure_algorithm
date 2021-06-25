package main

import (
	"fmt"
	"math/rand"
	"testing"
)
//测试1000次
//每次数据 1000 个
//时间复杂度 nlog(n)
//用时 (0.19s)(加上数据构造)
//java表现
//用时 (0.479s)(加上数据构造)
func TestInitArray(t *testing.T) {
	for i := 0; i < 1000; i++ {
		old_array := make([]int,0)
		for i := 0; i < 1000; i++ {
			old_array = append(old_array,int(rand.Int31n(10000)))
		}
		_, err := start(old_array)
		if err != nil {
			fmt.Printf("err=%#v\n", err)
			return
		}
		//fmt.Printf("array=%#v\n", array)
	}
}