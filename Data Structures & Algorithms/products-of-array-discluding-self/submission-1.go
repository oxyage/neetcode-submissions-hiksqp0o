func productExceptSelf(nums []int) []int {

	var result = make([]int, len(nums))
	// решение в лоб - пройти дважды по массиву, опираясь на базовый элемент nums[i]
	for i := 0; i < len(nums); i ++ {

		// в результате держим единицу, чтобы не менять произведение
		result[i] = 1

		for j := 0; j < len(nums); j ++ {
			// пропустим базовый элемент
			if i == j {
				continue
			}

			// домножаем результат
			result[i] = result[i] * nums[j]
		}

	}

	return result
}
