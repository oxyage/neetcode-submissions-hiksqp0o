func trap(height []int) int {
	// O(t, n)
	// O(s, 1)

	if len(height) == 0 {
		return 0
	}

	l, r := 0, len(height) - 1
	leftMax, rightMax := height[l], height[r]
	area := 0

	for l < r {
		if leftMax < rightMax {
			// двигаем вправо левый указатель, если он ниже правого
			l ++
			leftMax = max(leftMax, height[l])
			area += leftMax - height[l]
		} else {
			// иначе двигаем правый указатель влево
			r --
			rightMax = max(rightMax, height[r])
			area += rightMax - height[r]

		}

	}

	return area
}
