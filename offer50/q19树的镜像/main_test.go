package main

import (
	"fmt"
	"testing"
)
//功能测试
func TestFonction(t *testing.T) {
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
}
//鲁棒测试
func TestRobust(t *testing.T) {
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

}