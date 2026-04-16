// bruteforce решение
func countBits(n int) []int {
	res := make([]int, 0, n + 1)


	for i := 0; i <= n; i++ {
		res = append(res, calcBit(i))
	}

	
	return res
}


func calcBit(num int) int {

	c := 0
	// перевести в строку бит и подсчитать количество единиц
	for _, v := range fmt.Sprintf("%b", num) {
		if v == '1' {
			c++
		}
	}
	return c 
}