func dailyTemperatures(temperatures []int) []int {
	var result = make([]int, 0, len(temperatures))

// Return an array result where result[i] is the number of days after the ith day
// before a warmer temperature appears on a future day. 
// If there is no day in the future where a warmer temperature will appear for the ith day, set result[i] to 0 instead.

// result[i] - количество дней после i-ого (включая), прежде чем температура будет выше чем в i-ый день

	for i, today_t := range temperatures {
		a := 1
		var j int
		for j = i + 1; j < len(temperatures); j ++ {
			future_t := temperatures[j]

			if future_t > today_t {
				break
			}
			a++

		}

		if j == len(temperatures) {
			a = 0
		}
		
		result = append(result, a)
	}


	return result
}
