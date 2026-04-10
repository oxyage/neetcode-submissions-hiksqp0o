/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    return depth(root)
}


func depth(root *TreeNode) int{
	if root == nil {
		return 0
	}

	return 1 + max(depth(root.Left), depth(root.Right))
}
