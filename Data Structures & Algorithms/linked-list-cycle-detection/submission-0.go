/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	// способ быстрого и медленного указателя
	// медленный растет на +1
	// быстрый растет на +2
	// если они однажды станут равными - это значит список замкнут
    var fastP *ListNode
	var slowP *ListNode
	for {
		// если следующего указателя нет, или послеследующего нет, то список не замкнут
		if head.Next == nil || head.Next.Next == nil {
			return false
		}

		// если указатели сравнялись - значит замкнут
		if fastP == slowP {
			return true
		}

		// быстрому указателю ставми +2
		fastP = head.Next.Next
		// медленному +1
		slowP = head.Next	
	}
	return false
}
