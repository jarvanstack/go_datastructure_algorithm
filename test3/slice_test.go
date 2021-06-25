package main

import (
	"fmt"
	"testing"
	"unsafe"
)
//代码来源.
//https://zhuanlan.zhihu.com/p/323417660
//自制切片测试
func TestName(t *testing.T) {
	//初始化结构体
	var arr = [2]byte{0, 0}
	var s1 = struct {
		array unsafe.Pointer
		len   int
		cap   int
	}{unsafe.Pointer(&arr), 2, 2}
	//将结构指针转化为切片指针*[]byte
	s2 := (*[]byte)(unsafe.Pointer(&s1))
	fmt.Println("结构体转切片")
	fmt.Println(*s2)
	fmt.Println(arr)
	//通过切片修改底层数值
	(*s2)[0] = 0
	(*s2)[1] = 1
	fmt.Println("切片赋值")
	fmt.Println(*s2)
	fmt.Println(arr)
	//通过数组修改
	fmt.Println("数组赋值")
	arr[1] = 100
	fmt.Println(*s2) //切片赋值后
	fmt.Println(arr)
	//结构体指针和切片指针指向的对象是同一个地址
	fmt.Println("切片和结构体存储地址")
	fmt.Println(uintptr(unsafe.Pointer(&s1)))
	fmt.Println(uintptr(unsafe.Pointer(s2)))
}
func TestKuoRong(t *testing.T) {
	growslice()
}
func growslice() {
	//这种情况,经过了微调
	a := []byte{1, 2, 3} // len=3 ,cap = 3
	a = append(a, 4, 5)
	fmt.Println("a", len(a), cap(a)) //按照规则应该是2*old.cap=2*3=6,roundupsize(6*1)向上取整为8, 8/1=8,最终的cap=8,len=5
	//这种情况,恰好不用微调,符合扩容规则2,容量翻倍
	b := []int64{1, 2, 3} // len = 3,cap = 3
	b = append(b, 4, 5)
	fmt.Println("b", len(b), cap(b)) //按照规格 2*old.cap=2*3=6, roundupsize(6*8) = 48,满足取整,48/8=6,最终的cap=6,len=5
}
//做栈之前的测试
//栈的思路是，数据用切片存，初始化未 为 8 个值， 然后当做数组用，当最大值超过的时候用append，自动扩容
//2倍 16，实现动态扩容.
func TestStack(t *testing.T) {
	slice := []int{1,2,3,4,5,6,7, 8}
	fmt.Printf("len(slice)=%#v\n", len(slice))
	fmt.Printf("cap(slice)=%#v\n", cap(slice))
	slice = append(slice, 9)
	fmt.Printf("len(slice)=%#v\n", len(slice))
	fmt.Printf("cap(slice)=%#v\n", cap(slice))
}