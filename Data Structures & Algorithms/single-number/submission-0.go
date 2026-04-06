func singleNumber(nums []int) int {
	// можно использовать мапу?
	//  пройдем по элементам nums
	// будем записывать в мапу
	// если элемент уже есть - удалять из мапы
	// тот элемент, который останется в мапе - тому нет пары, его вернуть

	var set = make(map[int]struct{})
	for _, n := range nums {
		if _, ok := set[n];  ok {
			delete(set, n)
		} else {
			set[n] = struct{}{}
		}
	}


	for k, _ := range set {
		return k
	}
	return -1
}
