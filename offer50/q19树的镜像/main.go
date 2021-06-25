package main

import "fmt"

//题目：输入两棵二叉树A，输出二叉树的镜像树B。
//思路是前序遍历树，如果这个节点有左节点和右节点就交换他们.
//指针的话将会改变原来的二叉树A的结构.
//       8
//    8    7
// 9    2
//    4   7
type BTNode struct {
	data  int
	left  *BTNode
	right *BTNode
}

func mirrorTree(root *BTNode)  {
	if root == nil ||(root.left == nil &&root.right == nil) {
		return
	}
	//交换左子树和右子树的值
	tempNode := root.left
	root.left = root.right
	root.right = tempNode
	mirrorTree(root.left)
	mirrorTree(root.right)
}

func main() {
	//构造数据和测试.
	a1 := &BTNode{8, nil, nil}
	a2 := &BTNode{8, nil, nil}
	a3 := &BTNode{7, nil, nil}
	a4 := &BTNode{9, nil, nil}
	a5 := &BTNode{2, nil, nil}
	a6 := &BTNode{4, nil, nil}
	a7 := &BTNode{7, nil, nil}
	a1.left = a2
	a1.right = a3
	a2.left = a4
	a2.right = a5
	a5.left = a6
	a5.right = a7
	b1 := &BTNode{8, nil, nil}
	b2 := &BTNode{9, nil, nil}
	b3 := &BTNode{2, nil, nil}
	b1.left = b2
	b1.right = b3
	PrintBT(a1)
	fmt.Printf("\n")
	//PrintBT(b1)
	fmt.Printf("string=%s\n", "After mirrorTree")
	mirrorTree(a1)
	PrintBT(a1)
}

//先序遍历打印.
//遇到就打印.
//8 8 9 2 4 7 7
func PrintBT(root *BTNode) {
	if root == nil {
		return
	}
	fmt.Printf("%#v ", root.data)
	PrintBT(root.left)
	PrintBT(root.right)
}
