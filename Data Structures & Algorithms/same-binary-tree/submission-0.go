/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {

	// база рекурсии - когда заканчиваются все листы дерева
	if (p == nil && q == nil) {
		return true
	}

	// p | q

	// nil | nil == true - база рекурсии
	// nil | !nil == false - не равные листы
	// !nil | nil = false - не равные листы
	// !nil | !nil = isSameTree() || p.Val == q.Val - равные листы, рекурсивный запуск

	if (p == nil && q != nil || p != nil && q == nil || p.Val != q.Val){
		return false
	}


	// если один из результатов не true, деревья не равны
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
