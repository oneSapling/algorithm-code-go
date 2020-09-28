package main

// DFS深度搜索(分治的方法)
func DFS(root *TreeNode) []int{
	result := make([]int,0)
	if root == nil {
		return result;
	}
	left := DFS(root.Left)
	right := DFS(root.Right)
	// 合并操作
	result  = append(result,root.Val) //先序遍历的时候需要把值先放结果数组
	result =  append(result, left...)
	result = append(result,right...)
	result = result[:len(result)-1]
	return result
}