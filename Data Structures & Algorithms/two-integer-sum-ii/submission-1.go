func twoSum(numbers []int, target int) []int {

	// задача на два указателя
	// по массиву идем слева и справа
	// благодаря тому, что массив отсортирован 
	// правый указатель - грубая подстройка, поскольку справа лежат самые большие элементы
	// левый указатель - точная подстройка

	// ищем сумму двух чисел, которая будет равна target, начиная с элементов:
	// sum(numbers[left] , numbers[rigth]) <> target?,  где right = len(numbers) - 1

	// если sum < target, то нужна точная подстройка, сдвигаем левый индекс вправо
	// если sum > target, то нужна грубая подстройка, сдвигаем правый индекс влево
	// если sum = target, вернем индексы left, rigth


	// рекурсия или в цикле
	for left, right := 0, len(numbers) - 1; left <= right; {

		// текущие значения
		left_n, right_n := numbers[left], numbers[right]
		sum := left_n + right_n
		// выходим, если нашли индексы
		if sum == target {
			return []int{left + 1, right + 1}
		}

		// подстраиваем точно
		if sum < target {
			left += 1
			continue
		}

		// подстраиваем грубо
		if sum > target {
			right -= 1
			continue
		}
	}
	return []int{-1, -1}
}


