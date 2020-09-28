package main

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
