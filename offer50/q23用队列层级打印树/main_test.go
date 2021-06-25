package main

import "testing"
//基本测试
func Test_levelPrint(t *testing.T) {
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
//鲁棒性测试
func TestRobust(t *testing.T) {
	levelPrint(nil)

}
//鲁棒性测试
func TestRobust2(t *testing.T) {
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
	levelPrint(n2)
	levelPrint(n3)
	levelPrint(n4)
	levelPrint(n5)
	levelPrint(n6)
	levelPrint(n7)
}