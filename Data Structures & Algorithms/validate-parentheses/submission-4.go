func isValid(str string) bool {
    stack := Stack{}

	rules := map[rune]rune {
		'{':'}',
		'[':']',
		'(':')',
	}

	for _, r := range str {
		

		// открывающая скобка?
		if _, ok := rules[r]; ok {
			stack.Push(r)
		} else {
			// закрывающая скобка

			var lastEl rune
			var err error
	
			// проверим пустой элемен
			if lastEl, err = stack.Last(); err != nil {
				return false
			}

			// последний элемент стека имеет другую пару на закрытие
			if rules[lastEl] != r {
				return false
			}

			// пару нашли валидную, надо удалить из стека
			stack.Pop()
		}

	}

	return stack.IsEmpty()
}


// реализация стека
// методы: Push, Pop, Last (peak), Length
type Stack struct {
	items []rune
	length int
}


func(s *Stack) Len() int{
	return s.length
}

func(s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func(s *Stack) Push(v rune) {
	s.items = append(s.items, v)
	s.length += 1
}

func(s *Stack) Pop() {
	s.items = s.items[ : s.Len() - 1]
	s.length -= 1
}

func(s *Stack) Last() (rune, error) {
	if s.IsEmpty() {
		return ' ', fmt.Errorf("empty")
	}

	return s.items[ s.Len() - 1], nil
}




