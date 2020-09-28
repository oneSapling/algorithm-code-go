package main

// 后序遍历二叉树
func postOrder(root *TreeNode) []int{
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
