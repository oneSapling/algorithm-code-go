package main

import (
	"fmt"
)

//二叉树节点实现
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
测试方法
 */
func main() {
	// 创建一个二叉树
	TreeNodeHead := createBinaryTree()
	// result := preOrder(n1) //[1 2 4 6 7 8 3 5]
	// result := houOrderTree(n1)
	// result := fenzhi(n1) //[1 2 4 6 7 8 3 5]
	// 层次遍历
	result := DFS(TreeNodeHead)
	fmt.Println("遍历结果",result)
}

/*
创建一个二叉树，用来测试用
*/
func createBinaryTree() *TreeNode{
	n1 := &TreeNode{}
	n2 := &TreeNode{}
	n3 := &TreeNode{}
	n4 := &TreeNode{}
	n5 := &TreeNode{}
	n6 := &TreeNode{}
	n7 := &TreeNode{}
	n8 := &TreeNode{}
	n1.Val=1
	n2.Val=2
	n3.Val=3
	n4.Val=4
	n5.Val=5
	n6.Val=6
	n7.Val=7
	n8.Val=8
	n1.Left = n2
	n2.Left = n4
	n4.Right = n6
	n6.Left = n7
	n6.Right = n8
	n1.Right = n3
	n3.Right = n5
	return n1
}