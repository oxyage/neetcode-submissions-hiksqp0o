/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1 -> 2 -> 3 -> 4 - |
//      ^- - - - - - -|

// fastP = nil, slowP = nil
// fastP = 3, slowP = 2
// fastP = 2, slowP = 3
// fastP = 4, slowP = 4, return true


// 1 - > 2 -> 3 -> nil
// fastP = nil, slowP = nil
// fastP = 3, slowP = 2
// fastP = nil, slowP = 3 => return false

func hasCycle(head *ListNode) bool {
	// способ быстрого и медленного указателя
	// медленный растет на +1
	// быстрый растет на +2
	// если они однажды станут равными - это значит список замкнут
    var fastP *ListNode = head
	var slowP *ListNode = head

	for {

		if fastP == nil || slowP == nil {
			return false
		}

		if fastP.Next == nil || fastP.Next.Next == nil {
			return false
		}

		if slowP.Next == nil {
			return false
		}

		// быстрому указателю ставми +2
		fastP = fastP.Next.Next
		// медленному +1
		slowP = slowP.Next	

		// если указатели сравнялись - значит замкнут
		if slowP == fastP {
			return true
		}


	}
	return false
}
