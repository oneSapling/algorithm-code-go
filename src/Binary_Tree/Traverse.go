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
前中后遍历方法，
 */
func main() {
	// 创建一个二叉树
	TreeNodeHead := createBinaryTree()
	// result := preOrder(n1) //[1 2 4 6 7 8 3 5]
	// result := houOrderTree(n1)
	// result := fenzhi(n1) //[1 2 4 6 7 8 3 5]
	// 层次遍历
	result := levelOrder(TreeNodeHead)
	fmt.Println("遍历结果",result)
}
/**
前序遍历二叉树
 */
func preOrder(root *TreeNode) []int{
	result := make([]int, 0)
	if(root == nil){
		return result
	}
	stack := make([]*TreeNode,0)
	for root != nil || len(stack)>0 {
		for root!=nil {
			stack =append(stack, root)
			result =  append(result, root.Val)
			root = root.Left
		}
		//弹栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}
/**
中序遍历二叉树
*/
func midOrder(root *TreeNode) []int{
	result := make([]int, 0)
	if(root == nil){
		return result
	}
	stack := make([]*TreeNode,0)
	for root != nil || len(stack)>0 {
		for root!=nil {
			stack =append(stack, root)
			root = root.Left
		}
		//弹栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result =  append(result, node.Val)
		root = node.Right
	}
	return result
}
// 后序遍历二叉树
func houOrderTree(root *TreeNode) []int{
	if root== nil{
		return make([]int,0)
	}
	result := make([]int,0)
	stack := make([]*TreeNode,0)
	var lastView *TreeNode
	for root!=nil || len(stack)>0{
		for root != nil {
			stack =  append(stack, root)
			root = root.Left
		}
		node:= stack[len(stack)-1]
		if node.Right == nil || node.Right == lastView {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastView = node
		}else {
			root = node.Right
		}
	}
	return result
}
// DFS深度搜索(分治的方法)
func fenzhi(root *TreeNode) []int{
	result := make([]int,0)
	if root == nil {
		return result;
	}
	left := fenzhi(root.Left)
	right := fenzhi(root.Right)
	// 合并操作
	result  = append(result,root.Val) //先序遍历的时候需要把值先放结果数组
	result =  append(result, left...)
	result = append(result,right...)
	result = result[:len(result)-1]
	return result
}

//按层遍历搜索BFS
func levelOrder(root *TreeNode) [][]int{
	result := make([][]int,0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int,0)
		l := len(queue)
		for i:=0; i < l ;i++ {
			//出队-把当前层的节点都出对列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue,level.Left)
			}
			if level.Right != nil {
				queue = append(queue,level.Right)
			}
		}
		result = append(result, list)
	}
	return result
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