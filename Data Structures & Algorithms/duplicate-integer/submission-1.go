func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
	    for _, num := range nums {
		        seen[num] = struct{}{}
				    }
    return len(seen) < len(nums)
	
}
