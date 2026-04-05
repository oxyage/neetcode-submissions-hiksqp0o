func search(nums []int, target int) int {
	// первый запуск рекурсивной функции
	return recursiveSearch(0, len(nums) - 1, nums, target)
}

func recursiveSearch(left,right int, nums []int, target int) int {

	// база рекурсии = останавливаемся, если окно поиска закрылось
	if left > right {
		return -1
	}

	// текущий средний элемент - вычисляем от сдвига влево на половину всей длины
	middle := left + (right - left) / 2

	// если элемент по середине = target - заканчиваем рекурсию
	if nums[middle] == target {
		return middle
	}

	// если средний элемент меньше целевого
	// необходимо сдвинуть окно вправо, то есть левую границу до середины + 1
	if nums[middle] < target {
		return recursiveSearch( middle + 1 , right, nums, target )
	}

	// в другом случае, средний элемент больше целевого
	// необходимо сдвинуть окно влево, то есть правую границу до середины - 1
	return recursiveSearch(left, middle - 1, nums, target)
}
