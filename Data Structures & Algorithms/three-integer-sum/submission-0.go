

func threeSum(nums []int) [][]int {

	// задача тоже на два поинтера. слева и справа

	// пробовать сортировку?
	// [-1,0,1,2,-1,-4]
	// -4, -1, -1, 0, 1, 2

	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i ++ {

		a := nums[i]
		if a > 0 {
			// нет смысла идти i-ым индексом, если число положительное (потому что в сумме i, l, r никогда не дадут 0)
			break
		}

		// не повторяем предыдущий элемент - убираем дубликаты
		if i > 0 && a == nums[i - 1] {
			continue
		}

        l, r := i+1, len(nums)-1
		for l < r {
			// todo
			if a + nums[l] + nums[r] == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}


			if nums[l] + nums[r] + a < 0{
				l++
			}

			if nums[l] + nums[r] + a > 0{
				r --
			}

		}

	}

	return result
}
