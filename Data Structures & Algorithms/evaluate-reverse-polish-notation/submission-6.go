func evalRPN(tokens []string) int {
	// ["4","13","5","/","+"]
	// 4 + (13 / 5)

	//tokens=["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
	// 22

	// на стек кладем только числа
	// когда приходит оператор, выполняем действия с числами
	// возвращаем на стек результат

	var stack = make([]int, 0, len(tokens))
	
	for _, t := range tokens {
		
		// if got number, not operand
		if !( t=="+" || t=="-" || t=="*" || t=="/" ) {
			num, err := strconv.Atoi(t)
			if err != nil {
				continue
			}

			stack = append(stack, num)
			fmt.Printf("-- currentStack %v\n", stack)

			continue
		}


		stackLen := len(stack)
		a := stack[stackLen - 2]
		b := stack[stackLen - 1]
		// got operand
		result := calcTopStack(a, b, t)


		fmt.Printf("calcTopStack: %d %s %d = %d\n", a, t,b, result)

		stack = append(stack[0:stackLen - 2], result)

		fmt.Printf("-- currentStack %v\n", stack)


	}

	if len(stack) != 1 {
		return -999
	}

	return stack[0]
}

func calcTopStack(a, b int, operand string) int {

	switch operand 	{
		case "+": 
			return a + b
		case "-":
			return a - b
		case "/":
			return int(math.Floor(float64(a / b)))
		case "*":
			return a * b
	}
	
	return 0
}


