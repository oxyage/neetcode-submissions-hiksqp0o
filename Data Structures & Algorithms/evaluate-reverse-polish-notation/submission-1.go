func evalRPN(tokens []string) int {


	// на стек кладем только числа
	// когда приходит оператор, выполняем действия с числами
	// возвращаем на стек результат

	var stack = make([]int, 0, len(tokens))
	
	for _, t := range tokens {
		
		if t != "+" && t != "-" && t != "*" && t != "/" {
			// then number


			num, err := strconv.Atoi(t)
			if err != nil {
				continue
			}

			stack = append(stack, num)
			continue
		}

		// got operand
		result := calcStack(stack, t)
		fmt.Printf("calckStack: %v operand: %s, result: %d\n", stack, t, result)
		stack = []int{result}
	}

	if len(stack) != 1 {
		return -999
	}

	return stack[0]
}

func calcStack(stack []int, operand string) int {
	var result = 0
	if operand == "*" || operand == "/" {
		result = 1
	}

	for _, n := range stack {
		switch operand {
			case "+": 
				result += n
			case "-":
				result -= n
			case "/":
				result /= n
			case "*":
				result *= n
		}
	}
	
	return result
}


