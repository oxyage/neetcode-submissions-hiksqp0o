/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    
	// задача сводится к объединению двух связанных списков
	// аналогия с двумя указателями, которые идут слева и справа по связному списку
	// но здесь придется использовать метод медленного и быстрого указателей
	// медленный указатель идет по левой части, начиная с первого узла и идет +1
	// второй указатель начинает со второго узла и прыгает +2
	// т.о. медленный указатель остановится в конце первой половины
	// быстрый указатель будет в конце второй половины списка

	// 1. Найдем середину связанного списка fast, slow pointer
	slow, fast := head, head.Next
	// начинаем с head.Next != nil AND head.Next.Next != nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next // сдвигаем на один элемент медленный указатель
		fast = fast.Next.Next // на два сдвигаем быстрый указатель
	}

	// для четного списка SLOW указатель на середине списка, FAST в конце
	// 1 -> 2 -> 3 -> 4
	//      S         F  

	// для нечетного списка SLOW в середине, FAST указатель на nil
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> nil
	//                S              	   F


	// чтобы корректно объединить списки, необходимо развернуть правую часть
	// second начало правой части списка
	second := slow.Next // 
	var prev *ListNode
	slow.Next = nil // обнуляем конец левой части, потому что это будет конец нового списка
	for second != nil { 
		tmp := second.Next // 
		second.Next = prev // 
		prev = second // 
		second = tmp // 
	}

	// 0: 5 -> 6 -> 7 -> nil
	// 1: nil <- 5 х 6 -> 7 -> nil
	// 2: nil <- 5 <- 6 х 7 -> nil
	// 3: nil <- 5 <- 6 <- 7 х nil
	// 4: nil <- 5 <- 6 <- 7


	// теперь можно объединять два массива
	// prev будет находиться на последней позиции второй половины, 
	// то есть в начале второго связного списка

	left, right := head, prev
	for right != nil { // правая часть всегда будет равна либо короче левой
		leftNext, rightNext := left.Next, right.Next
		left.Next = right
		right.Next = leftNext
		left, right = leftNext, rightNext 

	}


}
