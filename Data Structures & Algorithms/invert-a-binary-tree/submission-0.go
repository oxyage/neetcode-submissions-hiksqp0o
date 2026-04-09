/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	// мне просто нравится делать отдельную функцию для рекурсии
	return recursiveInvert(root)
}

// на вход принимать только элемент, который считаешь корневым
func recursiveInvert(root *TreeNode) *TreeNode {
	// база рекурсии - если этого листа нет, дальше некуда идти
	if root == nil {
		return nil
	}

	// меняем местами так, потому что аргумент в функцию передан по указателю
	root.Left, root.Right = root.Right, root.Left

	// запускаем инверсию других узлов без возврата - по указателю и так всё инвертируется
	recursiveInvert(root.Left)
	recursiveInvert(root.Right)

	// возвращаем тот самый лист, который передан
	return root
}
