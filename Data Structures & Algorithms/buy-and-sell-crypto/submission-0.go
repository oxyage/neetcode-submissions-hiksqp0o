func maxProfit(prices []int) int {

	// вернуть из функции maxProfit
	var maxProfit int = 0
	var minPrice = prices[0]


	// prices=[7,6,4,3,1]
	// i=0, p=7, buy=7, profit=0, maxProfit=0
	// i=1, p=6, buy=6, profit=0, maxProfit=0
	// i=2, p=4, buy=4, profit=0, maxProfit=0
	// i=3, p=3, buy=3, profit=0, maxProfit=0
	// i=4, p=1, buy=1, 

	// prices=[10,1,5,6,7,1]
	/*
	i=0, p=10, buy=10, profit=0, maxProfit=0
	i=1, p=1, buy=1, profit=0, maxProfit=0
	i=2, p=5, buy=1, profit=4, maxProfit=4
	i=3, p=6, buy=1, profit=5, maxProfit=5
	i=4, p=7, buy=1, profit=6, maxProfit=6
	i=5, p=1, buy=1, profit=0, maxProfit=6
	*/
	
	// пойдем по всем ценам
	for _, todayPrice := range prices {

		// вычислим текущий профит?
		todayProfit := todayPrice - minPrice

		// покупаем ли сегодня?
		if todayPrice <= minPrice {
			minPrice = todayPrice
		} 		

		// текущий профит - это текущая цена минус минимальная цена
		if todayProfit > maxProfit {
			maxProfit = todayProfit
		}
	}


	return maxProfit
}
