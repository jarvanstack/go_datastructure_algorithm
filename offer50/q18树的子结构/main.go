package main

import "fmt"

//题目：输入两棵二叉树A和B，判断B是不是A的子结构。二叉树结点的定义如下：
//思路是找到头结点，然后依次遍历B的其他节点，是否匹配.
//Go没有现成的stack你就老实用递归.挺好.
//       8
//    8    7
// 9    2
//    4   7
type BTNode struct {
	data  int
	left  *BTNode
	right *BTNode
}

func IsSubTree(mainTreeRoot, subTreeRoot *BTNode) bool {
	if mainTreeRoot == nil || subTreeRoot == nil {
		return false
	}
	result := false
	//找到和subTree匹配的头结点.
	if mainTreeRoot.data == subTreeRoot.data {
		//子树匹配
		result = match(mainTreeRoot, subTreeRoot)
	}
	if result == false {
		//如果匹配失败，匹配左子树
		result = IsSubTree(mainTreeRoot.left, subTreeRoot)
	}
	if result == false {
		//如果匹配失败，匹配右子树
		result = IsSubTree(mainTreeRoot.right, subTreeRoot)
	}
	return result
}

//遍历子树，判断是否匹配.
func match(mainTreeNode, subTreeNode *BTNode) bool {
	//处理null
	if mainTreeNode == nil && subTreeNode != nil {
		return false
	}
	//如果子树完成遍历返回true
	if subTreeNode == nil {
		return true
	}
	if mainTreeNode.data != subTreeNode.data {
		return false
	}
	return match(mainTreeNode.left,subTreeNode.left) && match(mainTreeNode.right,subTreeNode.right)
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
	PrintBT(b1)
	fmt.Printf("IsSubTree(a1,b1)=%#v\n", IsSubTree(a1, b1))
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
