func largestRectangleArea(heights []int) int {
	// todo: разобрать
	n := len(heights)
	stack := make([]int, 0)

	leftMost := make([]int, n)
	for i := range leftMost {
		leftMost[i] = -1
	}

	for i := 0; i < n; i ++ {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) > 0 {
			leftMost[i] = stack[len(stack) - 1]
		}

		stack = append(stack, i)
	}

	stack = stack[:0]
	rightMost := make([]int, n)
	for i := range rightMost {
		rightMost[i] = n
	}

	for i := n - 1; i >= 0; i -- {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) > 0 {
			rightMost[i] = stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		leftMost[i] ++
		rightMost[i] --
		area := heights[i] * (rightMost[i] - leftMost[i] + 1)
		maxArea = max(area, maxArea)

	}

	return maxArea
}
