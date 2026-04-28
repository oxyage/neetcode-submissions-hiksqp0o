func maxProfit(prices []int) int {

	// два указателя, l = buy, r = right
	l, r := 0, 1
	maxProfit := 0

	// идем правым указателем до конца массива
	for r < len(prices) {
		// проверяем, возможен ли профит?
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			// обновляем максимальный профит, если был
			maxProfit = max(profit, maxProfit)
		} else {
			// если нет профита, то сдвигаем левый указатель дальше, можно перепрыгнуть на следующий r
			l = r
		}
		r += 1
	}


	return maxProfit
}
