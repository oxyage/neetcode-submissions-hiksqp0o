func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers) - 1
	result := make([]int, 0, 2)
	for l < r {
		// поскольку слайс отсортирован, движение левого указателя - точная подстройка
		// движение правого указателя влево - грубая подстройка

		// пусть res - промежуточный результат, res = numbers[l] + numbers[r]
		res := numbers[l] + numbers[r]
		if res == target{
			return []int{numbers[l], numbers[r]}
		}
		if res > target {
		// тогда, если res > target, то нужна грубая подстройка, двигаем правый указатель влево
			r--
		} else {
		// иначе двигаем левый указатель вправо
			l ++
		}
	}


	return result
}
