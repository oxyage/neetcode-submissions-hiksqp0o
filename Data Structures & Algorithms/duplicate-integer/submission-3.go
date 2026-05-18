func hasDuplicate(nums []int) bool {
    var numSet = make(map[int]int, len(nums))
    for _, val := range nums {
        if _, ok := numSet[val]; ok {
            return true
        } else {
            numSet[val] = 1
        }
    }
    return false
}
