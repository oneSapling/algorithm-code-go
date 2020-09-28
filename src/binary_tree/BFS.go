package main

//按层遍历搜索BFS
func BFS(root *TreeNode) [][]int{
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
