package main

import (
	"math/rand"
	"testing"
	"time"
)

//题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的
//比如 126 345 合并之后就是 123456.
//合并2个拍好序的链表，
//方案1,一个一个移动，维护2个指针即可.
//方案2,可以滑动方式的移动，维护4个指针.
//功能测试
func TestFunction(t *testing.T) {
	//创建一个有序的链表
	head1 := InitLinked([]int{1,2,6})
	//打印.
	//PrintLinked(head1)
	//创建一个有序的链表
	head2 := InitLinked([]int{3,4,5,5})
	//打印.
	//PrintLinked(head2)
	//合并链表,
	linked := MergeLinked(head1, head2)
	PrintLinked(linked)
}
//功能测试2
func TestFunction2(t *testing.T) {
	//创建一个有序的链表
	head1 := InitLinked([]int{1,2,6})
	//打印.
	//PrintLinked(head1)
	//创建一个有序的链表
	head2 := InitLinked([]int{3,4,4,4,5})
	//打印.
	//PrintLinked(head2)
	//合并链表,
	linked := MergeLinked2(head1, head2)
	PrintLinked(linked)
	t.Logf("测试通过\n")
}
//鲁棒性测试
func TestRobust(t *testing.T) {
	MergeLinked(nil,nil)
	MergeLinked(InitLinked([]int{1}),InitLinked([]int{1}))
	MergeLinked(InitLinked([]int{1}),nil)
	MergeLinked(nil,InitLinked([]int{1}))
	t.Logf("测试成功")
}
//压力测试
//为什么测试失败？？
func TestPressure(t *testing.T) {
	var data1 [100_0000]int
	var data2 [100_0000]int
	for i := 0; i < 100_0000; i++ {
		data1[i] = 10 * i + (rand.Int()%10)
		data2[i] = 10 * i + (rand.Int()%10)
	}
	head1 := InitLinked(data1[:])
	head2 := InitLinked(data2[:])
	start := time.Now()
	for i := 0; i < 1000; i++ {
		MergeLinked(head1,head2)
		MergeLinked(head1,head2)
	}
	end := time.Now()
	t.Logf("测试通过,用时%v",(end.UnixNano() -start.UnixNano()))
}