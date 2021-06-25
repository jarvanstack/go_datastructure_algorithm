package main

import (
	"DataStructureAlgorithm/utils"
	"fmt"
)

//题目：一层一层的打印二叉树，比如
//       1
//    2    3
// 4    5 6    7
// 打印顺序为 1234567，
//使用队列报错每一次层的节点即可，十分简单，之前我也做过，重点是，队列，如果 go 没有队列，我就
//自己做一个

//树的节点.
type treeNode struct {
	data int
	left *treeNode
	right *treeNode
}

func main() {
	n1 := &treeNode{data: 1}
	n2 := &treeNode{data: 2}
	n3 := &treeNode{data: 3}
	n4 := &treeNode{data: 4}
	n5 := &treeNode{data: 5}
	n6 := &treeNode{data: 6}
	n7 := &treeNode{data: 7}
	n2.left = n4
	n2.right = n5
	n3.left = n6
	n3.right = n7
	n1.left = n2
	n1.right  = n3
	levelPrint(n1)
}
//层顺序遍历
func levelPrint(root *treeNode)  {
	//处理空.
	if root == nil {
		fmt.Printf("%s\n", "root不能为空")
		return
	}
	//
	queue := utils.NewQueue()
	err := queue.Push(root)
	if err != nil {
		fmt.Printf("err=%#v\n", err)
		return
	}
	//遍历直到队列中位空
	//输出队列中一个元素就将这个元素Pop，然后在将子树(如果有)入队列.
	for !queue.IsEmpty() {
		pop := queue.Pop().(*treeNode)
		fmt.Printf("pop=%#v\n", pop.data)
		//入子树入对了
		if pop.left!=nil {
			_ = queue.Push(pop.left)
		}
		if pop.right!=nil {
			_ = queue.Push(pop.right)
		}
	}
}