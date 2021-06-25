package main

import (
	"DataStructureAlgorithm/utils"
	"errors"
	"fmt"
	"math"
	"reflect"
)
//使用辅助 minStack 储存每次最小值和每次pop之后的下一个最小值.
//怎么做到？
//辅助栈栈顶保留最小的值，如果push的值更小，就保留更小的值，如果push的值不够小，就保留之前的栈顶的最小的值.
//包含最小值的栈.
type MinStack struct {
	stack *utils.Stack
	minStack *utils.Stack
}
//构造函数
func NewMinStack()*MinStack  {
	return &MinStack{stack: utils.NewStack(),minStack: utils.NewStack()}
}
//如果push的值更小，就保留更小的值，
//如果push的值不够小，就保留之前的栈顶的最小的值.
func (this *MinStack) Push(data int)error  {
	//如果是第一个值就不用比较了
	if this.minStack.IsEmpty() {
		_ = this.stack.Push(data)
		_ = this.minStack.Push(data)
		return nil
	}
	//如果不是第一个就需要和之之前的比较。
	//拿到最小栈的最小栈顶值.
	//取出
	peek, _ := this.minStack.Peek()
	//如果类型不匹配返回错误
	if reflect.TypeOf(peek).Kind() != reflect.Int {
		return errors.New("数据类型类型不匹配,不是int")
	}
	//如果小于最小值就新插入最小值.
	if data < peek.(int) {
		_ = this.stack.Push(data)
		_ = this.minStack.Push(data)
		return  nil
	}
	//如果大于最小值，就依旧保留之前的最下值在 minStack的top
	_ = this.stack.Push(data)
	_ = this.minStack.Push(peek.(int))
	return  nil
}
//IsEmpty
func (this *MinStack)IsEmpty()bool  {
	return this.minStack.IsEmpty() || this.stack.IsEmpty()
}
//Pop
func (this *MinStack) Pop()(int,error)  {
	//栈为空
	if this.IsEmpty() {
		return 0,errors.New("栈为空")
	}
	//栈不为空。
	_, _ = this.minStack.Pop()
	pop, _ := this.stack.Pop()
	return pop.(int),nil
}
//拿到最小的值.就peek minStack
func (this *MinStack) GetMin() (int,error) {
	if this.IsEmpty() {
		return math.MinInt32,errors.New("栈为空")
	}
	peek, err := this.minStack.Peek()
	if err != nil {
		return 0,errors.Unwrap(err)
	}
	return peek.(int), nil
}
func main() {
	minStack := NewMinStack()
	fmt.Printf("minStack.IsEmpty()=%#v\n", minStack.IsEmpty())
	_ = minStack.Push(3)
	min, _ := minStack.GetMin()
	fmt.Printf("minStack.GetMin()=%#v\n",min )
	_ = minStack.Push(4)
	min, _ = minStack.GetMin()
	fmt.Printf("minStack.GetMin()=%#v\n",min )
	_ = minStack.Push(2)
	min, _ = minStack.GetMin()
	fmt.Printf("minStack.GetMin()=%#v\n",min )
	_ = minStack.Push(1)
	fmt.Printf("minStack.IsEmpty()=%#v\n", minStack.IsEmpty())
	min, _ = minStack.GetMin()
	fmt.Printf("minStack.GetMin()=%#v\n",min )
	_, _ = minStack.Pop()
	_, _ = minStack.Pop()
	min, _ = minStack.GetMin()
	fmt.Printf("minStack.GetMin()=%#v\n",min )

}
