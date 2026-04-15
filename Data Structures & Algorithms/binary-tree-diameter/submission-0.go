/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {

	maxDiameter := 0
	recDiameter(root, &maxDiameter)
	return maxDiameter
}


// рекурсивная функция должна возвращать высоту поддерева, 
// чтобы родительский узел корретктно посчитал свой вклад в общий путь
func recDiameter(root *TreeNode, diameter *int) int {

	// база рекурсии
	if (root == nil) {
		return 0
	}

	// вычисляем диаметр через левого потомка
	leftPath := recDiameter(root.Left, diameter)
	// и правого потомка
	rightPath := recDiameter(root.Right, diameter)

	// складываем и сравниваем с максимальным диаметром снаружи
	currentD := leftPath + rightPath
	if currentD > *diameter {
		*diameter = currentD
	}



	return 1 + max(leftPath, rightPath)
}

