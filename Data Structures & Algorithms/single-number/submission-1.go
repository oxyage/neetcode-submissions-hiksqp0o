func singleNumber(nums []int) int {
	// задача может быть решена через битовые операции

	// 3 = 11
	// 2 = 10

	result := 0
	for _, num := range nums {
		result = result ^ num
	}

	return result

	// 3 XOR 3 = 0

	// 6 = 110
	// 7 = 111
	// 8 = 1000

}
