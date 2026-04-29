func longestConsecutive(nums []int) int {

	// преобразование в мапу для доступа O(1)
 	numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

	longest := 0
	// проход циклом по мапе
	for num := range numSet {
		if _, left_exists := numSet[num - 1]; !left_exists {
			// если левого соседа num не существует, значит это новая последовательность

			length := 1
			// начинаем новый цикл, чтобы посчитать длину вправо от num
			for {

				// если у num есть правый сосед, увеличиваем длину
				if _, right_exists := numSet[num + length]; right_exists {
					length += 1
				} else {
					// справа num соседа нет
					break
				}

				// выбираем более длинную последовательность
				longest = max(longest, length)
			}
		}
	}
	return longest
}
