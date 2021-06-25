package utils

import (
	"fmt"
	"testing"
)
//100w 0.08 s
//1000w (1.16s)
//使用copy替代append 1000w  (0.62s)
//使用*2替代25% append 1000w  (0.36s)
func TestPress(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 10_000; i++ {
		_ = stack.Push(0)
	}
	fmt.Printf("stack.cap=%#v\n", stack.cap)
	fmt.Printf("stack.topIndex=%#v\n", stack.topIndex)
	fmt.Printf("len(stack.data)=%#v\n", len(stack.data))
	fmt.Printf("cap(stack.data)=%#v\n", cap(stack.data))
}
//1.push 自己做的栈的测试.
func TestFunction1(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 9; i++ {
		_ = stack.Push(i)
	}
	fmt.Printf("stack.size()=%#v\n", stack.size())
	fmt.Printf("stack.cap=%#v\n", stack.cap)
	fmt.Printf("stack.topIndex=%#v\n", stack.topIndex)
}
//为什么stack.cap为0
func TestCap(t *testing.T) {
	stack := NewStack()
	_ = stack.Push(1)
	fmt.Printf("stack.size()=%#v\n", stack.size())
	fmt.Printf("stack.cap=%#v\n", stack.cap)
}//因为扩容条件写错了
//切片的最大长度是int大小，所以...我们length还是使用int罢了.
//2.pop()继续测试 pop
func TestFunction2(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 9; i++ {
		_ = stack.Push(i)
	}
	pop, _ := stack.Pop()
	fmt.Printf("pop=%#v\n", pop)
	fmt.Printf("stack.size()=%#v\n", stack.size())
	fmt.Printf("stack.cap=%#v\n", stack.cap)
}
//3.size()
func TestFunction3(t *testing.T) {

	stack := NewStack()
	fmt.Printf("stack.isEmpty()=%#v\n", stack.IsEmpty())
	for i := 0; i < 8; i++ {
		_ = stack.Push(i)
	}
	pop, _ := stack.Pop()
	fmt.Printf("stack.size()=%#v\n", stack.size())
	fmt.Printf("stack.isEmpty()=%#v\n", stack.IsEmpty())
	fmt.Printf("pop=%#v\n", pop)
	fmt.Printf("stack.size()=%#v\n", stack.size())
	fmt.Printf("stack.cap=%#v\n", stack.cap)
}
//100w 0.08 s
//1000w (1.16s)
//使用copy替代append 1000w  (0.62s)
//使用*2替代25% append 1000w  (0.36s)
func TestMax(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 10_000; i++ {
		_ = stack.Push(0)
	}
	fmt.Printf("stack.cap=%#v\n", stack.cap)
	fmt.Printf("stack.topIndex=%#v\n", stack.topIndex)
	fmt.Printf("len(stack.data)=%#v\n", len(stack.data))
	fmt.Printf("cap(stack.data)=%#v\n", cap(stack.data))
}
func Test8(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 8; i++ {
		_ = stack.Push(i)
	}
	_ = stack.Push(8)
	_ = stack.Push(9)
}
//切片虽然动态扩缩容了，但是...
//length = 9 的切片下标也只能访问到 8 ,而不是 16 -1 = 15 懂了么，
//那我们就不用切片了，自己用指针做!!!试试. 自己写一个切片.
//我们有2个思路，一个是数组，如果超过下标就，新一个数组，然后 copy()
//2.想了解下切片的 unsafe.pointer怎么弄的.
//算了用第一种方式罢
//我们用数组试试.会不会快点.
func Test9(t *testing.T) {
	var temp [8<<1]int
	fmt.Printf("len(temp)=%#v\n", len(temp))
}
//topIndex 除了点问题
//99999 没有问题.
//超过 1024 就出了点问题
//删除1024还是这样，奇怪
//debug
//stack.cap=1033428
//stack.topIndex=999999
//len(stack.data)=1033428l
//cap(stack.data)=1033428
func Test9999(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 99_999999; i++ {
		_ = stack.Push(0)
	}
	fmt.Printf("stack.cap=%#v\n", stack.cap)
	fmt.Printf("stack.topIndex=%#v\n", stack.topIndex)
	fmt.Printf("len(stack.data)=%#v\n", len(stack.data))
	fmt.Printf("cap(stack.data)=%#v\n", cap(stack.data))
}
//栈pop测试
func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	_ = stack.Push(1)
	_ = stack.Push(2)
	peek, _ := stack.Peek()
	fmt.Printf("stack.Peek()=%#v\n", peek)
	_, _ = stack.Pop()
	peek, _ = stack.Peek()
	fmt.Printf("stack.Peek()=%#v\n", peek)
}