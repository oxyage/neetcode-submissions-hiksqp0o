func twoSum(nums []int, target int) []int {
    
	var set = make(map[int]int)
	for i, n := range nums {
		set[n] = i
	}

	for i, n := range nums {

		diff := target - n
		if j, ok := set[diff]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
