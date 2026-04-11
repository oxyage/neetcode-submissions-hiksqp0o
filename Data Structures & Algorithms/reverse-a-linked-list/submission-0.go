/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func reverseList(head *ListNode) *ListNode {
    

	// direct
	// 1 -> 2 -> 3 -> null
	// reversed
	// 3 -> 2 -> 1 -> null


 /**
 
 reverseList(1): head.Val=1, head.Next=2
	new_head = reverseList(2):
		reverseList(2): head.Val=2, head.Next=3
			new_head = reverseList(3):
				reverseList(3): head.Val=3, !head.Next=nil!
				return 3 
			
	
reverseList(3): head.Val=3, head.Next=nil

  
 */

	// база рекурсии - следующий узел - конец, или у его следующего узла конец
	if head == nil || head.Next == nil{
		return head
	}

	// запускаем рекурсию для следующего узла
	new_head := reverseList(head.Next)


	head.Next.Next = head
	head.Next = nil	


	return new_head
}

