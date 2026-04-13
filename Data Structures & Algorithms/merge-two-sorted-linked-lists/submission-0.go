/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // нужен какой-то фиктивный узел, чтобы идти по итерациям нового списка
    head := &ListNode{
        Val: -1111,
        Next: nil,
    }
    // текущий узел начинается с фиктивного - он будет двигаться по новому списку
    var current *ListNode 

    // заходим сюда только в случае, если оба списка не пустые
    // как только один из списков станет пустым, мы выйдем из цикла
    
    // обычный for цикл, но в качестве итератора - указатель на текущий узел
    for current = head; list1 != nil && list2 != nil; current = current.Next {

        // надо выбрать что добавлять в current = меньший элемент
        if list1.Val <= list2.Val {
            // элемент из первого списка меньше, поэтому в currentNext добавляем его
            current.Next = list1
            // сдвигаем в самом списке указатель на следующий элемент
            list1 = list1.Next
        } else {
            // мЕньший элемент во втором списке - сдвигаем его указатель
            current.Next = list2
            list2 = list2.Next
        }
    }
    // если вышли из цикла, значит один из списков пуст
    // осталось докинуть в текущий список остаток
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

    return head.Next
}
