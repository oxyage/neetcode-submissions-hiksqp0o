// найти заданную сумму элементов
func twoSum(nums []int, target int) []int {
    // необходимо найти первую комбинацию элементов, которые дают сумму = target

	// двумя циклами пойдем по массиву nums
	// i считаем базовым - первое слагаемое
	// j второе слагаемое, начинаем с j = i + 1
	// если nums[i] + nums[j] == target, выходим с [i, j]
	// пара i,j всегда есть (по условию)
	for i, val_i := range nums {
		for j := i+1; j < len(nums); j ++ {
			val_j := nums[j]
			if val_i + val_j == target {
				return []int{i, j}
		 	}
		}
	}
	return []int{0, 1}
}
