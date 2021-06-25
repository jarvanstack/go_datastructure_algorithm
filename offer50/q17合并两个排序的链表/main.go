package main

import (
	"fmt"
)

//题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的
//比如 126 345 合并之后就是 123456.
//合并2个拍好序的链表，
//方案1,一个一个移动，维护2个指针即可.
//方案2,可以滑动方式的移动，维护4个指针.
type Node struct {
	data int
	next *Node
}

func main() {
	//创建一个有序的链表
	head1 := InitLinked([]int{1,2,6})
	//打印.
	//PrintLinked(head1)
	//创建一个有序的链表
	head2 := InitLinked([]int{3,4,5})
	//打印.
	//PrintLinked(head2)
	//合并链表,
	linked := MergeLinked(head1, head2)
	PrintLinked(linked)
}
//至此每次多个节点merge的方法.
func MergeLinked2(head1,head2 *Node)*Node  {
	if head1 == nil && head2 == nil {
		//如果2个值都为空就返回空
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	//每次挑选2个链表里面比较小的，放到结果链表里面
	resultNode := &Node{}
	l1 := head1
	r1 := head1
	l2 := head2
	r2 := head2
	//初始化首位值.
	if  l1.data <= l2.data{
		resultNode = l1
		l1 = l1.next
		//r 和 l 是并行的.
		r1 = l1
	}else {
		resultNode = l2;
		l2 = l2.next
		//r 和 l 是并行的.
		r2 = l2
	}
	resultNodeTemp := resultNode
	//当 l1 和 l2都是空就退出循环
	for  true{
		if l1 == nil && l2 == nil {
			break
		}
		//移动滑块
		//移动剩下的l2
		if l1 == nil  {
			resultNodeTemp.next = l2
			break
		}else if l2 == nil {
			//移动剩下的l1
			resultNodeTemp.next = l1
			break
		}
		//移动滑块1.
		//
		if l1.data < l2.data {
			//r1返回的时候必然不能为空.
			for  {
				if r1.next == nil || r1.next.data > l2.data {
					break
				}
				r1 = r1.next
			}
			resultNodeTemp.next = l1
			resultNodeTemp = r1
			//r1 = l1 = 6.
			r1=r1.next
			l1 = r1
		}else {
			//移动滑块2 345
			//r1返回的时候必然不能为空.
			for  {
				if r2.next == nil || r2.next.data > l1.data {
					break
				}
				r2 = r2.next
			}
			resultNodeTemp.next = l2
			resultNodeTemp = r2
			//l2 = r2 = null
			r2=r2.next
			l2 = r2
		}

	}
	return resultNode
}
//merge tow sorted linked  by head1 and head2.Return new linked head node.
func MergeLinked(head1,head2 *Node)*Node  {
	if head1 == nil && head2 == nil {
		//如果2个值都为空就返回空
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	//每次挑选2个链表里面比较小的，放到结果链表里面
	resultNode := &Node{}
	temp1 := head1
	temp2 := head2
	//初始化首位值.
	if  temp1.data <= temp2.data{
		resultNode = temp1
		temp1 = temp1.next
	}else {
		resultNode = temp2;
		temp2 = temp2.next
	}
	resultNodeTemp := resultNode
	for  true{
		if temp1==nil && temp2 == nil {
			//遍历完就退出.
			break
		}
		//如果2为空或者1比较小等.
		//我们就，将1插入到结果集中。并且将结果集下移动
		if temp2==nil || temp1.data <= temp2.data{
			resultNodeTemp.next = temp1
			temp1 = temp1.next
			resultNodeTemp = resultNodeTemp.next
		}else{
			resultNodeTemp.next = temp2
			temp2 = temp2.next
			resultNodeTemp = resultNodeTemp.next
		}
	}
	return resultNode
}
//init node with n length.
func InitLinked(array []int)*Node  {
	head := &Node{data: array[0], next: nil}
	temp := head
	for i := 1; i < len(array); i++ {
		node := &Node{data: array[i], next: nil}
		temp.next = node
		temp = temp.next
	}
	return head
}
//Print the list
func PrintLinked(head *Node)  {
	temp := head
	for true {
		if temp==nil {
			break
		}
		fmt.Printf("node.data=%#v\n", temp.data)
		temp = temp.next
	}
}
