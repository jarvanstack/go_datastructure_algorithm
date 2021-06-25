package main

import (
	"fmt"
	"testing"
)

func PrintMinStack(stack MinStack) {
	fmt.Printf("stack.stack=%#v\n", stack.stack)
	fmt.Printf("stack.minStack=%#v\n", stack.minStack)
}
func TestMinStack_Pop(t *testing.T) {
	stack := NewMinStack()
	_ = stack.Push(3)
	_ = stack.Push(4)
	_ = stack.Push(2)
	_ = stack.Push(1)
	fmt.Printf("stack.stack=%#v\n", stack.stack)
	fmt.Printf("stack.minStack=%#v\n", stack.minStack)
	_, _ = stack.Pop()
	fmt.Printf("stack.stack=%#v\n", stack.stack)
	fmt.Printf("stack.minStack=%#v\n", stack.minStack)
	//测试minStack pop失败，直接无效.我们试试打印pop错误
	pop, err := stack.Pop()
	if err != nil {
		fmt.Printf("err=%#v\n", err)
	}
	fmt.Printf("pop=%#v\n", pop)
	//栈是空的，为什么呢，我们试试
	fmt.Printf("stack.IsEmpty()=%#v\n", stack.IsEmpty())
	//这玩意就是空的.我们试试二级
	fmt.Printf("stack.stack=%#v\n", stack.stack)
	fmt.Printf("stack.stack.IsEmpty()=%#v\n", stack.stack.IsEmpty())
}
//测试getMain
func TestMinStack_GetMin(t *testing.T) {
	stack := NewMinStack()
	_ = stack.Push(3)
	_ = stack.Push(4)
	_ = stack.Push(2)
	_ = stack.Push(1)
	min, _ := stack.GetMin()
	fmt.Printf("min=%#v\n", min)
	_, _ = stack.Pop()
	min, _ = stack.GetMin()
	fmt.Printf("min=%#v\n", min)
	_, _ = stack.Pop()
	min, _ = stack.GetMin()
	fmt.Printf("min=%#v\n", min)
	_, _ = stack.Pop()
	min, _ = stack.GetMin()
	fmt.Printf("min=%#v\n", min)
	_, _ = stack.Pop()
	min, _ = stack.GetMin()
	fmt.Printf("min=%#v\n", min)
}
