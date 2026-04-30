func maxArea(heights []int) int {
	maxArea := 0
	l, r := 0, len(heights) - 1
	for l < r {
		d := r - l // длина
		h := min(heights[l], heights[r])
		maxArea = max(maxArea, d * h)
		
		if heights[l] < heights[r] {
			l += 1
		} else if heights[l] > heights[r] {
			r -= 1
		} else {
			l += 1
		}
		
	}

	return maxArea
}
