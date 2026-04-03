func isValid(s string) bool {
   // если заявлена симметричность не заявлена, необходимо использовать стек

	// даже несимметричная строка будет содержать четное количество символов
	if len(s) % 2 != 0 {
		return false
	}

// составим карту соответствий
	var rules = map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	// будем хранить открытые скобки
	var stack = make([]byte, 0, len(s))


	// проходим по строке, смотрим каждую скобку
	for _, parenthes := range s {

		v := byte (parenthes)

		if _, ok := rules[v]; ok {
			// открывающая скобка, кладем в стек
			stack = append(stack, v)
		} else {
			// закрывающая скобка, пытаемся удалить из стека

			if len(stack) == 0 {
				return false
			}

			lastStack := stack[len(stack) - 1] 

			if rules[lastStack] != v {
				return false
			} 

			stack = stack [:len(stack) - 1]
		}

	}


	// на выходе стек должен быть пустым
	return len(stack) == 0
}
