package main

import (
	"DataStructureAlgorithm/utils"
	"fmt"
)

//题目：输入两个整数序列，第一个序列表示栈的压入顺序，
//请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
//例如序列1、2、3、4、5是某栈的压栈序列，序列4、5、3、2、1是该压栈序列对应的一个弹出序列，
//但4、3、5、1、2就不可能是该压栈序列的弹出序列。
//example
//push {1,2,3,4,5} pop {4,5,3,2,3}
//第一个希望被pop的是4,于是需要4是栈顶
//于是一次压入 1,2,3,4
//此时栈顶是 4,弹出 4,pop 下一个是 5
//继续依次入栈直到栈顶是 5
// 1 2 3 5
//此时栈顶是 5 弹出 5, pop 下一个是 3
//此时栈顶是 3 弹出 3,pop 下一个是 2
//此时栈顶是 2 弹出 2,pop 下一个是 1
//此时栈顶是 1 弹出 1,pop 下一个结束了，所以返回 true
//while 循环的条件是
//如果 push数组已经走完，并且栈中为空或者栈顶不匹配pop
//example2
//push {1,2,3,4,5} pop {4,5,2,1,3}
//第一个希望pop的是. 4 于是 入栈 1，
//栈顶不是 4 ,继续入栈 2
//栈顶不是 4 ,继续入栈 3
//栈顶不是 4 ,继续入栈 4
//栈顶是 4 ,4 出栈,出栈轮到 5
//栈顶不是 5，5 入栈,
//栈顶不是 5,5 出栈 ,出栈轮到2
//栈顶不是2，push 已经全部入栈,退出循环.

//更具压栈的顺序判断是否为弹出的顺序
func IsPopList(pushList, popList []int) bool {
	//非法输入
	if pushList == nil || popList == nil {
		return false
	}
	pushLength := len(pushList)
	popLength := len(popList)
	if pushLength <= 0 || popLength <= 0 || popLength != pushLength {
		return false
	}
	//判断
	stack := utils.NewStack()
	//推入第一个元素.
	_ = stack.Push(pushList[0])
	pushIndex := 1
	popIndex := 0
	for true {
		if popIndex == popLength {
			return true
		}
		for pushIndex < pushLength && stack.Peek().(int) != popList[popIndex] {
			_ = stack.Push(pushList[pushIndex])
			pushIndex++

		}
		if pushIndex == pushLength && stack.Peek().(int) != popList[popIndex] {
			 return false
		}
		for !stack.IsEmpty() && popIndex < popLength && stack.Peek().(int) == popList[popIndex] {
			_ = stack.Pop()
			popIndex++
		}

	}
	return false
}
func main() {
	fmt.Printf("IsPopList()=%#v\n", IsPopList([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}))
	fmt.Printf("IsPopList()=%#v\n", IsPopList([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Printf("IsPopList()=%#v\n", IsPopList([]int{1, 2, 3, 4, 5}, []int{4, 5, 2, 3, 1}))
	fmt.Printf("IsPopList()=%#v\n", IsPopList([]int{1}, []int{1}))
	fmt.Printf("IsPopList()=%#v\n", IsPopList(nil, []int{1}))
	fmt.Printf("IsPopList()=%#v\n", IsPopList(nil, nil))
}
